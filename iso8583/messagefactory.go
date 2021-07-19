package iso8583

import (
	lib8583 "github.com/moov-io/iso8583"
)

// MessageFactory to create instances of the ISO8583 messages
type MessageFactory struct {
	Specification lib8583.MessageSpec
}

// DefaultMessageFactory - Creates a default instance of the MessageFactory
func DefaultMessageFactory(specFile string) *MessageFactory {
	// TODO: implement it to load it from an external file
	return &MessageFactory{Specification: *lib8583.Spec87}
}

// NewInstance - Creates a new instance
func (mf *MessageFactory) NewInstance(mti string) *lib8583.Message {
	isoMsg := lib8583.NewMessage(&mf.Specification)
	isoMsg.MTI(mti)
	return isoMsg
}

// DefaultInstance - Creates a default instance
func (mf *MessageFactory) DefaultInstance() *lib8583.Message {
	isoMsg := lib8583.NewMessage(&mf.Specification)
	return isoMsg
}
