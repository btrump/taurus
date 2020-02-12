package ttt

import (
	"errors"
	"fmt"
	"log"

	"github.com/btrump/taurus-server/pkg/message"
)

// Engine is a container for the state of the game world
type Engine struct {
	state State
}

// New returns a new, initialized state machine
func New() *Engine {
	log.Printf("fsm::New(): Creating new Engine")
	e := &Engine{}
	e.initialize()
	return e
}

// PlayerCurrent returns the currently active player
func (f Engine) PlayerCurrent() string {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()
	// return f.State.Players[f.State.Order[f.State.TurnCounter%len(f.State.Order)]]
	return f.state.Players[f.state.Order[f.state.TurnCounter%len(f.state.Order)]].ID
}

// PlayerAdd adds a player to the list of players and the order list
func (e *Engine) PlayerAdd(n string) (message.Response, error) {
	if len(e.GetPlayers()) > 1 {
		msg := "ttt::PlayerAdd(): Fauiled to add player %s"
		return message.NewResponse(false, fmt.Sprintf(msg, n)), errors.New(msg)
	}
	i := len(e.GetPlayers())
	e.state.Players[i] = NewPlayer(string(i), n)
	e.state.Order[i] = i
	return message.NewResponse(true, fmt.Sprintf("ttt::PlayerAdd(): Added player %s", n)), nil
}

func (e *Engine) GetPlayers() [2]*Player {
	return e.state.Players
}

// GetState returns the current engine state
func (e *Engine) GetState() interface{} {
	return &e.state
}

// Stats returns statistics about the engine
func (e *Engine) Stats() interface{} {
	return struct{}{}
}
