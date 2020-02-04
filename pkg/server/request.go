package server

import (
	"errors"
	"log"
	"time"

	"github.com/btrump/taurus-server/internal/message"
	"github.com/btrump/taurus-server/pkg/phase"
	"github.com/rs/xid"
)

func (s *Server) ReceiveRequest(req message.Request) message.Response {
	log.Printf("server::receiveRequest(): Got message with command '%s' from user '%s'", req.Command, req.UserID)
	res := s.evaluateMessage(req)
	req.ID = xid.New().String()
	res.ID = xid.New().String()
	s.Messages = append(s.Messages, req, res)
	return res
}

func (s *Server) requestValidate(m message.Request) error {
	var e error
	switch m.Command {
	case "GAME_START":
		if s.State.Phase != phase.PHASE_PRE {
			e = errors.New("server::requestValidate(): Could not start game")
		}
	case "GAME_END":
		if s.State.Phase == phase.PHASE_COMPLETED {
			e = errors.New("server::requestValidate(): Could not end game. Game already ended")
		}
	case "TURN_END":
		if false {
			e = errors.New("server::requestValidate(): Could not end turn")
		}
	case "NEXT_PHASE":
		if s.State.Phase != phase.PHASE_STARTED {
			e = errors.New("server::requestValidate(): Could not advance phase. Not in STARTED state")
		}
	default:
		e = errors.New("server::requestValidate(): Did not recognize command")
	}
	return e
}

func (s *Server) requestHandle(m message.Request) message.Response {
	var msg string
	switch m.Command {
	case "GAME_START":
		s.State.Phase = phase.PHASE_STARTED
	case "GAME_END":
		s.State.Phase = phase.PHASE_COMPLETED
	case "TURN_END":
		s.State.TurnCounter++
		msg = "Ending turn."
		if s.State.TurnCounter%len(s.State.Order) == 0 {
			msg += " Ending round."
			s.State.RoundCounter++
		}
	case "NEXT_PHASE":
		s.State.Phase++
	default:
		msg = "server::evaluateMessage(): Did not recognize command. This should never happen, because request was already validated"
	}
	return message.Response{
		Timestamp: time.Now(),
		Success:   true,
		Message:   msg,
	}
}
