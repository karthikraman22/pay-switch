package imf

import "time"

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
