package iso8583

import (
	"encoding/binary"

	"github.com/panjf2000/gnet"
)

type iso8583Codec struct {
	*gnet.LengthFieldBasedFrameCodec
}

func buildIso8583Codec(fieldLength int) *iso8583Codec {
	encoderConfig := gnet.EncoderConfig{
		ByteOrder:                       binary.BigEndian,
		LengthFieldLength:               fieldLength,
		LengthAdjustment:                0,
		LengthIncludesLengthFieldLength: false,
	}
	decoderConfig := gnet.DecoderConfig{
		ByteOrder:           binary.BigEndian,
		LengthFieldOffset:   0,
		LengthFieldLength:   fieldLength,
		LengthAdjustment:    0,
		InitialBytesToStrip: fieldLength,
	}
	return &iso8583Codec{gnet.NewLengthFieldBasedFrameCodec(encoderConfig, decoderConfig)}
}
