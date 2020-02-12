package engine

import "github.com/btrump/taurus-server/pkg/message"

type State interface{}

type Engine interface {
	GetState() interface{}
	Stats() interface{}
	PlayerAdd(string) (message.Response, error)
	PlayerCurrent() string
	Validate(message.Request) (message.Response, error)
	Execute(message.Request) (message.Response, error)
}
