/*
Package fsm provides a state machine that manages state of the game world
*/
package fsm

import (
	"fmt"
	"log"

	"github.com/btrump/taurus-server/pkg/message"
)

// Phase is the current phase of the game state
type Phase int

// Valid phases
const (
	PRE Phase = iota
	STARTED
	COMPLETED
)

// FSM is a container for the state of the game world
type FSM struct {
	State State
}

// initialize sets the initial values of the state machine
func (f *FSM) initialize() {
	log.Printf("fsm::initialize(): Initializing FSM")
	f.State = State{
		Phase: PRE,
	}
	f.State.Players = make(map[string]*Player)
}

// New returns a new, initialized state machine
func New() *FSM {
	log.Printf("fsm::New(): Creating new FSM")
	f := FSM{}
	f.initialize()
	return &f
}

// PlayerCurrent returns the currently active player
func (f *FSM) PlayerCurrent() *Player {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()
	return f.State.Players[f.State.Order[f.State.TurnCounter%len(f.State.Order)]]
}

// PlayerAdd adds a player to the list of players and the order list
func (f *FSM) PlayerAdd(id string, n string) message.Response {
	f.State.Players[id] = NewPlayer(id, n)
	f.State.Order = append(f.State.Order, id)
	return message.NewResponse(true, fmt.Sprintf("fms::PlayerAdd(): Added player %s", id))
}
