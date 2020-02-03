package server

import (
	"errors"
	"log"
	"time"

	"github.com/btrump/taurus/internal/message"
	"github.com/btrump/taurus/pkg/client"
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

type state struct {
	Order        []string
	RoundCounter int
	TurnCounter  int
	Phase        phase
}

type phase int

const (
	PHASE_PRE phase = iota
	PHASE_STARTED
	PHASE_COMPLETED
)

var Config serverConfig
var Clients []clientConnection
var Messages []interface{}
var Chat []string
var State state

func initialize() {
	log.Printf("server::initialize(): Initializing")
	State = state{
		Phase: PHASE_PRE,
	}
}
func configure(config []map[string]string) {
	log.Printf("server::configureServer(): PLACEHOLDER")
	Config = serverConfig{
		ID:            "uniqueid",
		Port:          8081,
		Name:          "taurus-server",
		ServerVersion: "development",
		APIVersion:    "development",
		Started:       time.Now(),
	}
}

func connectClient(client client.Client) {
	log.Printf("server::connectClient(): %s connected", client.ID)
	log.Printf("server::connectClient(): %s appended to client list", client.ID)
	Clients = append(Clients, clientConnection{
		Client:    &client,
		Connected: time.Now(),
	})
	log.Printf("server::connectClient(): %s appended to order list", client.ID)
	State.Order = append(State.Order, client.ID)
}

func receiveRequest(req message.Request) message.Response {
	log.Printf("server::receiveRequest(): Got message with command '%s' from user '%s'", req.Command, req.UserID)
	res := evaluateMessage(req)
	req.ID = xid.New().String()
	res.ID = xid.New().String()
	Messages = append(Messages, req, res)
	return res
}

func validateRequest(m message.Request) error {
	var e error
	switch m.Command {
	case "GAME_START":
		if State.Phase == PHASE_PRE {
			State.Phase = PHASE_STARTED
			// e = errors.New("server::evaluateMessage(): Starting the game")
			// success = true
		} else {
			e = errors.New("server::evaluateMessage(): Could not start game")
		}
	case "GAME_END":
		if State.Phase != PHASE_COMPLETED {
			State.Phase = PHASE_COMPLETED
			// e = errors.New("server::evaluateMessage(): Ending the game")
			// success = true
		} else {
			e = errors.New("server::evaluateMessage(): Could not end game")
		}
	case "TURN_END":
	case "NEXT_PHASE":
		if State.Phase != PHASE_COMPLETED {
			State.Phase++
			// e = errors.New("server::evaluateMessage(): Advancing to next phase")
			// success = true
		} else {
			e = errors.New("server::evaluateMessage(): Could not advance phase")
		}
	default:
		e = errors.New("server::evaluateMessage(): Did not recognize command")
	}
	return e
}

func evaluateMessage(m message.Request) message.Response {
	if err := validateRequest(m); err != nil {
		log.Printf("server::evaluateMessage(): validateRequest failure")
		return message.Response{
			Timestamp: time.Now(),
			Success:   false,
			Message:   err.Error(),
		}
	}
	return srvHandleRequest(m)
}

func srvHandleRequest(m message.Request) message.Response {
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
		State.TurnCounter++
		msg = "Ending turn."
		if State.TurnCounter%len(State.Order) == 0 {
			msg += " Ending round."
			State.RoundCounter++
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

// Start accepts a configuration KVP object, and starts both the game engine and API server
func Start(config ...map[string]string) {
	initialize()
	configure(config)
	startAPI()
}
