package iso8583

import (
	lib8583 "github.com/moov-io/iso8583"
)

type Handler interface {
	Handle(in *lib8583.Message) (*lib8583.Message, error)
}

func GetHandler(mti string, mf *MessageFactory) (Handler, error) {
	var handler Handler

	switch mti {
	case "0200":
		handler = &Handler02xx{mf: mf}
	}
	return handler, nil
}
