package imf

import "time"

type SignOnRq struct {
	Version       string // "admn.001.001.01"
	MsgId         string
	DateTime      time.Time
	FunctionCode  string // "signon"
	CorrelationId string
	ParticipantId string
}

type SignOnRs struct {
	Version       string // "admn.002.001.01"
	MsgId         string
	DateTime      time.Time
	FunctionCode  string // "signon"
	CorrelationId string
	ParticipantId string
	Status        string // "ACTC - Accepted technical validation"
}
