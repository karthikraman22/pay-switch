package iso8583

import (
	lib8583 "github.com/moov-io/iso8583"
)

type Handler02xx struct {
	mf *MessageFactory
}

func (h Handler02xx) Handle(in *lib8583.Message) (*lib8583.Message, error) {
	outMsg := h.mf.NewInstance("0210")
	outMsg.Field(2, "4242424242424242")
	outMsg.Field(3, "123456")
	outMsg.Field(4, "100")
	return outMsg, nil
}
