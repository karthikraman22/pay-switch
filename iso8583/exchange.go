package iso8583

import (
	"context"

	lib8583 "github.com/moov-io/iso8583"
	"github.com/panjf2000/gnet"
)

type Exchange struct {
	In      *lib8583.Message
	Out     *lib8583.Message
	Ctx     context.Context
	ReplyTo gnet.Conn
	MF      *MessageFactory
	Meta    map[string]string
}

func BuildExchange(in *lib8583.Message, ctx context.Context, conn gnet.Conn, mf *MessageFactory) *Exchange {
	return &Exchange{In: in, Ctx: ctx, ReplyTo: conn, MF: mf, Meta: make(map[string]string)}
}
