package ep

import (
	"context"
	"fmt"

	"achuala.in/payswitch/iso8583"
	"go.uber.org/zap"
)

type EndpointCfg struct {
	Id         string
	Desc       string
	Type       string
	HostNPort  string
	Proto      string
	Tls        bool
	AuthScheme string
}

type Endpoint interface {
	Init() error
	Run()
	Stop(context.Context) error
}

type EndpointMgr struct {
	endpoints map[string]Endpoint
	logger    *zap.Logger
}

func NewEndpointMgr(logger *zap.Logger) *EndpointMgr {
	return &EndpointMgr{endpoints: map[string]Endpoint{}, logger: logger}
}

func (epm *EndpointMgr) Shutdown(ctx context.Context) {
	for _, e := range epm.endpoints {
		e.Stop(ctx)
	}
}

func (epm *EndpointMgr) NewServerEndpoint(cfg *EndpointCfg) (*Endpoint, error) {
	if cfg.Proto == "tcp/8583" {
		ep := iso8583.NewIso8583Server(fmt.Sprintf("tcp://%s", cfg.HostNPort), epm.logger)
		err := ep.Init()
		if err != nil {
			return nil, err
		}
		ep.Run()
		epm.endpoints[cfg.Id] = ep
	}
	return nil, nil
}

func (epm *EndpointMgr) NewClientEndpoint(cfg *EndpointCfg) (*Endpoint, error) {
	return nil, nil
}
