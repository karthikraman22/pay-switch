package imf

import (
	"time"

	"achuala.in/payswitch/generated/payhub/imx"
	lib8583 "github.com/moov-io/iso8583"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EchoRq struct {
	Version       string // "admn.005.001.01"
	MsgId         string
	DateTime      time.Time
	FunctionCode  string // "echo"
	CorrelationId string
	ParticipantId string
}

type EchoRs struct {
	Version       string // "admn.006.001.01"
	MsgId         string
	DateTime      time.Time
	FunctionCode  string // "echo"
	CorrelationId string
	ParticipantId string
	Status        string // "ACTC - Accepted technical validation"
}

func Iso8583ToImfEchoRq(isoMsg *lib8583.Message) (*imx.MxAdm00500101, error) {
	id, _ := isoMsg.GetString(2)
	time, _ := isoMsg.GetString(3)
	ts, _ = time.Parse("MMDDhhmmss", time)
	hdr := &imx.MxHdr{MsgId: id, CreatedDateTime: timestamppb.New(ts)}
	return &imx.MxAdm00500101{Header: hdr}, nil
}
