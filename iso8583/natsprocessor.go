package iso8583

import (
	"achuala.in/payswitch/auth"
	"achuala.in/payswitch/core"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

type NatsProessor struct {
	authenticator auth.Authenticator
	nc            *nats.Conn
}

func NewNatsProcessor() *NatsProessor {
	return &NatsProessor{authenticator: &auth.WhiteListAuthenticator{}}
}

func (dh *NatsProessor) Process(exchange *Exchange) error {
	client, err := core.LoadClientByDeviceId(uuid.NewString())
	if err != nil {

		return err
	}
	err = dh.authenticator.Authenticate(client)
	if err == nil {
		mti, err := exchange.In.GetMTI()
		handler, err := GetHandler(mti, exchange.MF)
		out, err := handler.Handle(exchange.In)
		exchange.Out = out
		return err
	}
	return err
}
