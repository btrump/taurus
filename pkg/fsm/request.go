package fsm

import (
	"errors"
	"strconv"

	"github.com/btrump/taurus-server/pkg/message"
)

// Validate ensures that a request is valid
func (f *FSM) Validate(m message.Request) (message.Response, error) {
	var err error
	switch m.Command {
	case "GAME_START":
		if f.State.Phase != PRE {
			err = errors.New("server::Validate(): Could not start game")
		}
	case "GAME_END":
		if f.State.Phase == COMPLETED {
			err = errors.New("server::Validate(): Could not end game. Game already ended")
		}
	case "TURN_END":
		if false {
			err = errors.New("server::Validate(): Could not end turn")
		}
	case "NEXT_PHASE":
		if f.State.Phase != STARTED {
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
func (f *FSM) Execute(m message.Request) (message.Response, error) {
	var err error
	var msg string
	switch m.Command {
	case "GAME_START":
		f.State.Phase = STARTED
		msg = "server::requestExecute(): Game started"
	case "GAME_END":
		f.State.Phase = COMPLETED
		msg = "server::requestExecute(): Game ended"
	case "TURN_END":
		if !f.isTurn(m.UserID) {
			msg = "server::requestExecute(): Not player's turn"
			err = errors.New(msg)
			break
		}
		f.State.TurnCounter++
		msg = "server::requestExecute(): Ending turn"
		if f.State.TurnCounter%len(f.State.Order) == 0 {
			msg += ". Ending round."
			f.State.RoundCounter++
		}
	case "NEXT_PHASE":
		f.State.Phase++
		msg = "server::requestExecute(): Advancing to next phase"
	case "MARK_TILE":
		tile, _ := strconv.ParseInt(m.Message, 0, 64)
		f.State.Data.Env[tile] = m.UserID
	default:
		msg = "server::requestExecute(): Did not recognize command. This should never happen, because request was already validated"
	}
	return message.NewResponse(err == nil, msg), err
}
