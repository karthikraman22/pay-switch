package imf

import "time"

type SignOffRq struct {
	Version       string // "admn.003.001.01"
	MsgId         string
	DateTime      time.Time
	FunctionCode  string // "signoff"
	CorrelationId string
	ParticipantId string
}

type SignOffRs struct {
	Version       string // "admn.004.001.01"
	MsgId         string
	DateTime      time.Time
	FunctionCode  string // "signoff"
	CorrelationId string
	ParticipantId string
	Status        string // "ACTC - Accepted technical validation"
}
