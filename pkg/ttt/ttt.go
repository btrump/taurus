/*
Package ttt provides a state machine that manages state of the game world
*/
package ttt

import (
	"errors"
	"log"
	"strconv"

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

// initialize sets the initial values of the state machine
func (e *Engine) initialize() {
	log.Printf("fsm::initialize(): Initializing Engine")
	e.state = State{
		Phase: PRE,
	}
	// e.state.Players = make(map[string]*Player)
	for i := range e.state.Data.Env {
		e.state.Data.Env[i] = "-"
	}
}

// Validate ensures that a request is valid
func (f *Engine) Validate(m message.Request) (message.Response, error) {
	var err error
	switch m.Command {
	case "GAME_START":
		if f.state.Phase != PRE {
			err = errors.New("server::Validate(): Could not start game")
		}
	case "GAME_END":
		if f.state.Phase == COMPLETED {
			err = errors.New("server::Validate(): Could not end game. Game already ended")
		}
	case "TURN_END":
		if false {
			err = errors.New("server::Validate(): Could not end turn")
		}
	case "NEXT_PHASE":
		if f.state.Phase != STARTED {
			err = errors.New("server::Validate(): Could not advance phase. Not in STARTED state")
		}
	case "MARK_TILE":
	default:
		err = errors.New("server::Validate(): Did not recognize command")
	}
	if err != nil {
		return message.NewResponse(false, err.Error()), err
	}
	return message.NewResponse(true, "server::Validate(): Valid command"), err
}

// Execute performs the command indicated by a request
func (f *Engine) Execute(m message.Request) (message.Response, error) {
	var err error
	var msg string
	switch m.Command {
	case "GAME_START":
		f.state.Phase = STARTED
		msg = "server::requestExecute(): Game started"
	case "GAME_END":
		f.state.Phase = COMPLETED
		msg = "server::requestExecute(): Game ended"
	case "TURN_END":
		if !f.isTurn(m.UserID) {
			msg = "server::requestExecute(): Not player's turn"
			err = errors.New(msg)
			break
		}
		f.state.TurnCounter++
		msg = "server::requestExecute(): Ending turn"
		if f.state.TurnCounter%len(f.state.Order) == 0 {
			msg += ". Ending round."
			f.state.RoundCounter++
		}
	case "NEXT_PHASE":
		f.state.Phase++
		msg = "server::requestExecute(): Advancing to next phase"
	case "MARK_TILE":
		tile, _ := strconv.ParseInt(m.Message, 0, 64)
		f.state.Data.Env[tile] = m.UserID
	default:
		msg = "server::requestExecute(): Did not recognize command. This should never happen, because request was already validated"
	}
	return message.NewResponse(err == nil, msg), err
}

func (e *Engine) GetScore(i int) int {
	return e.state.Score[i]
}
