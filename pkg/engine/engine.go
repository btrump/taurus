package engine

import "github.com/btrump/taurus-server/pkg/message"

// Engine is an interface that defines the contract for interaction with a server
type Engine interface {
	GetState() interface{}
	Stats() interface{}
	PlayerAdd(string) (message.Response, error)
	PlayerCurrent() string
	Validate(message.Request) (message.Response, error)
	Execute(message.Request) (message.Response, error)
}
