package core

import (
	"net"

	"github.com/google/uuid"
)

// Client is ...
type Client struct {
	Id         string
	DeviceId   string
	MerchantId string
	Addr       net.Addr
	// Key configuration if required
}

func LoadClientByDeviceId(deviceId string) (*Client, error) {
	return &Client{Id: uuid.NewString(), DeviceId: deviceId, MerchantId: uuid.NewString()}, nil
}
