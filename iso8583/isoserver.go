package iso8583

import (
	"context"
	"time"

	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pool/goroutine"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Iso8583Server struct {
	*gnet.EventServer
	Addr       string
	Codec      *iso8583Codec
	mf         *MessageFactory
	p          Processor
	workerPool *goroutine.Pool
	logger     *zap.Logger
}

func NewIso8583Server(addr string, logger *zap.Logger) *Iso8583Server {
	// TODO: Make the defaults configurable
	return &Iso8583Server{Addr: addr, Codec: buildIso8583Codec(2), mf: DefaultMessageFactory("iso8583-87.yaml"),
		p:          NewDefaultProcessor(),
		workerPool: goroutine.Default(),
		logger:     logger}
}

func (s *Iso8583Server) Init() error {
	return nil
}

func (s *Iso8583Server) Run() {
	go func() {
		if err := gnet.Serve(s, s.Addr,
			gnet.WithMulticore(true),
			gnet.WithCodec(s.Codec),
			gnet.WithTCPKeepAlive(time.Second*30),
			gnet.WithSocketRecvBuffer(8*1024),
			gnet.WithSocketSendBuffer(8*1024),
			gnet.WithReusePort(true),
			gnet.WithLogger(s.logger.Sugar()),
			gnet.WithLogLevel(zapcore.DebugLevel)); err != nil {
			s.logger.Fatal("iso8583 server start failed", zap.Error(err))
		}
	}()
}

func (s *Iso8583Server) Stop(ctx context.Context) error {
	return gnet.Stop(ctx, s.Addr)
}

func (s *Iso8583Server) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	s.logger.Info("iso8583 server started", zap.String("listen-address", srv.Addr.String()), zap.Bool("multi-cores", srv.Multicore), zap.Int("event-loops", srv.NumEventLoop))
	// TODO: if required register for echo messages, Check the config and register
	return
}

func (s *Iso8583Server) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	isoMsg := s.mf.DefaultInstance()
	if err := isoMsg.Unpack(frame); err == nil {
		s.logger.Debug("incoming message", zap.Any("payload", isoMsg))
		exchange := BuildExchange(isoMsg, nil, c, s.mf)
		s.p.Process(exchange)
		if outIsoMsg, err := exchange.Out.Pack(); err == nil && outIsoMsg != nil {
			go func() {
				if err := c.AsyncWrite(outIsoMsg); err != nil {
					s.logger.Error("error in writing", zap.Error(err))
				}
			}()
		}
	} else {
		s.logger.Error("invalid message received", zap.Error(err))
	}
	return
}

func (s *Iso8583Server) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {

	s.logger.Info("new client connected", zap.Any("local", c.LocalAddr()), zap.Any("remote", c.RemoteAddr()))
	return
}

func (s *Iso8583Server) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	s.logger.Info("client disconnected", zap.Error(err))
	s.workerPool.Release()
	return
}
