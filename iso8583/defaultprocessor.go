package iso8583

import (
	"achuala.in/pay-switch/auth"
	"achuala.in/pay-switch/core"
	"github.com/google/uuid"
)

// Processor is
type Processor interface {
	// Process
	Process(exchange *Exchange) error
}

type DefaultProcessor struct {
	authenticator auth.Authenticator
}

func NewDefaultProcessor() *DefaultProcessor {
	return &DefaultProcessor{authenticator: &auth.WhiteListAuthenticator{}}
}

func (dh *DefaultProcessor) Process(exchange *Exchange) error {
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
