package auth

import (
	"achuala.in/payswitch/core"
)

// Authenticator is
type Authenticator interface {
	// Authenticates the client
	Authenticate(client *core.Client) error
}

type WhiteListAuthenticator struct {
	approvedAddresses []string
}

func (wa *WhiteListAuthenticator) Authenticate(client *core.Client) error {
	return nil
}

func BuildAuthenticator() Authenticator {
	whiteListedIps := []string{"localhost", "127.0.0.1byte"}
	authenticator := &WhiteListAuthenticator{approvedAddresses: whiteListedIps}
	return authenticator
}
