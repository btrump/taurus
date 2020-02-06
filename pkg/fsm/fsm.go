/*
Package fsm provides a state machine that manages state of the game world
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

// State is a container for the objects in the game world
type State struct {
	Order        []string
	RoundCounter int
	TurnCounter  int
	Phase        Phase
}

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
}

// New returns a new, initialized state machine
func New() *FSM {
	log.Printf("fsm::New(): Creating new FSM")
	f := FSM{}
	f.initialize()
	return &f
}
