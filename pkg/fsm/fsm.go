/*
Package fsm provides a state machine that manages state
*/
package fsm

import "log"

// Phase is the current phase of the game state
type Phase int

// Valid phases
const (
	PRE Phase = iota
	STARTED
	COMPLETED
)

type State struct {
	Order        []string
	RoundCounter int
	TurnCounter  int
	Phase        Phase
}

type FSM struct {
	State State
}

func (f *FSM) initialize() {
	log.Printf("fsm::initialize(): Initializing FSM")
	f.State = State{
		Phase: PRE,
	}
}

func New() *FSM {
	log.Printf("fsm::New(): Creating new FSM")
	f := FSM{}
	f.initialize()
	return &f
}
