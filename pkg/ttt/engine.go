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
func (e Engine) PlayerCurrent() string {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()
	return e.state.Players[e.state.Order[e.state.TurnCounter%len(e.state.Order)]].ID
}

// PlayerAdd adds a player to the list of players and the order list
func (e *Engine) PlayerAdd(n string) (message.Response, error) {
	l := len(e.GetPlayers())
	if l > 1 {
		msg := "ttt::PlayerAdd(): Fauiled to add player %s"
		return message.NewResponse(false, fmt.Sprintf(msg, n)), errors.New(msg)
	}
	e.state.Players = append(e.state.Players, NewPlayer(n, n))
	e.state.Order = append(e.state.Order, l)
	return message.NewResponse(true, fmt.Sprintf("ttt::PlayerAdd(): Added player %s", n)), nil
}

func (e *Engine) GetPlayers() []*Player {
	return e.state.Players
}

// GetState returns the current engine state
func (e *Engine) GetState() interface{} {
	return e.state
}

// Stats returns statistics about the engine
func (e *Engine) Stats() interface{} {
	return struct{}{}
}

func (e *Engine) SetScore(p int, s int) int {
	e.state.Score[p] = s
	return s
}

// GetScore returns the score for player i
func (e *Engine) GetScore(i int) int {
	return e.state.Score[i]
}

func (e *Engine) SetPhase(p Phase) {
	e.state.Phase = p
}
