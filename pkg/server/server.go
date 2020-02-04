package server

import (
	"errors"
	"log"
	"time"

	"github.com/btrump/taurus-server/internal/message"
	"github.com/btrump/taurus-server/pkg/client"
	"github.com/btrump/taurus-server/pkg/phase"
	"github.com/btrump/taurus-server/pkg/state"
	"github.com/rs/xid"
)

type serverConfig struct {
	Started       time.Time
	ID            string
	Port          int
	Name          string
	ServerVersion string
	APIVersion    string
}

type clientConnection struct {
	Connected time.Time
	ID        string
	Client    *client.Client
}

type Server struct {
	Config   serverConfig
	Clients  []clientConnection
	Messages []interface{}
	Chat     []string
	State    state.State
}

func initialize(s *Server) {
	log.Printf("server::initialize(): Initializing a new server")
	s.State = state.State{
		Phase: phase.PHASE_PRE,
	}
}

func configure(s *Server, config []map[string]string) {
	log.Printf("server::configureServer(): PLACEHOLDER")
	s.Config = serverConfig{
		ID:            "uniqueid",
		Port:          8081,
		Name:          "taurus-server",
		ServerVersion: "development",
		APIVersion:    "development",
		Started:       time.Now(),
	}
}

func (s *Server) ClientConnect(client client.Client) {
	log.Printf("server::ClientConnect(): %s connected", client.ID)
	log.Printf("server::ClientConnect(): %s appended to client list", client.ID)
	s.Clients = append(s.Clients, clientConnection{
		Client:    &client,
		Connected: time.Now(),
	})
	log.Printf("server::ClientConnect(): %s appended to order list", client.ID)
	s.State.Order = append(s.State.Order, client.ID)
}

func (s *Server) ReceiveRequest(req message.Request) message.Response {
	log.Printf("server::receiveRequest(): Got message with command '%s' from user '%s'", req.Command, req.UserID)
	res := s.evaluateMessage(req)
	req.ID = xid.New().String()
	res.ID = xid.New().String()
	s.Messages = append(s.Messages, req, res)
	return res
}

func (s *Server) validateRequest(m message.Request) error {
	var e error
	switch m.Command {
	case "GAME_START":
		if s.State.Phase == phase.PHASE_PRE {
			s.State.Phase = phase.PHASE_STARTED
			// e = errors.New("server::evaluateMessage(): Starting the game")
		} else {
			e = errors.New("server::evaluateMessage(): Could not start game")
		}
	case "GAME_END":
		if s.State.Phase != phase.PHASE_COMPLETED {
			s.State.Phase = phase.PHASE_COMPLETED
			// e = errors.New("server::evaluateMessage(): Ending the game")
		} else {
			e = errors.New("server::evaluateMessage(): Could not end game")
		}
	case "TURN_END":
	case "NEXT_PHASE":
		if s.State.Phase != phase.PHASE_COMPLETED {
			s.State.Phase++
			// e = errors.New("server::evaluateMessage(): Advancing to next phase")
		} else {
			e = errors.New("server::evaluateMessage(): Could not advance phase")
		}
	default:
		e = errors.New("server::evaluateMessage(): Did not recognize command")
	}
	return e
}

func (s *Server) evaluateMessage(m message.Request) message.Response {
	if err := s.validateRequest(m); err != nil {
		log.Printf("server::evaluateMessage(): validateRequest failure")
		return message.Response{
			Timestamp: time.Now(),
			Success:   false,
			Message:   err.Error(),
		}
	}
	return s.requestHandle(m)
}

func (s *Server) requestHandle(m message.Request) message.Response {
	var msg string
	switch m.Command {
	case "GAME_START":
		// if State.Phase == PHASE_PRE {
		// 	State.Phase = PHASE_STARTED
		// 	// e = errors.New("server::evaluateMessage(): Starting the game")
		// 	// success = true
		// } else {
		// 	e = errors.New("server::evaluateMessage(): Could not start game")
		// }
	case "GAME_END":
		// if State.Phase != PHASE_COMPLETED {
		// 	State.Phase = PHASE_COMPLETED
		// 	// e = errors.New("server::evaluateMessage(): Ending the game")
		// 	// success = true
		// } else {
		// 	e = errors.New("server::evaluateMessage(): Could not end game")
		// }
	case "TURN_END":
		s.State.TurnCounter++
		msg = "Ending turn."
		if s.State.TurnCounter%len(s.State.Order) == 0 {
			msg += " Ending round."
			s.State.RoundCounter++
		}
	case "NEXT_PHASE":
		// if State.Phase != PHASE_COMPLETED {
		// 	State.Phase++
		// 	// e = errors.New("server::evaluateMessage(): Advancing to next phase")
		// 	// success = true
		// } else {
		// 	e = errors.New("server::evaluateMessage(): Could not advance phase")
		// }
	default:
		msg = "server::evaluateMessage(): Did not recognize command. This should never happen"
	}
	return message.Response{
		Timestamp: time.Now(),
		Success:   true,
		Message:   msg,
	}
}

// New accepts a configuration KVP object, and starts both the game engine and API server
func New(config ...map[string]string) Server {
	s := Server{}
	initialize(&s)
	configure(&s, config)
	return s
}
