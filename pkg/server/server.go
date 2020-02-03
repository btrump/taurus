package server

import (
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
	Order []string
	Turn  int
	Phase phase
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

func evaluateMessage(m message.Request) message.Response {
	var msg string
	success := false
	switch m.Command {
	case "GAME_START":
		if State.Phase == PHASE_PRE {
			State.Phase = PHASE_STARTED
			msg = "server::evaluateMessage(): Starting the game"
			success = true
		} else {
			msg = "server::evaluateMessage(): Could not start game"
		}
	case "GAME_END":
		if State.Phase != PHASE_COMPLETED {
			State.Phase = PHASE_COMPLETED
			msg = "server::evaluateMessage(): Ending the game"
			success = true
		} else {
			msg = "server::evaluateMessage(): Could not end game"
		}
	case "NEXT_PHASE":
		if State.Phase != PHASE_COMPLETED {
			State.Phase++
			msg = "server::evaluateMessage(): Advancing to next phase"
			success = true
		} else {
			msg = "server::evaluateMessage(): Could not advance phase"
		}
	default:
		msg = "server::evaluateMessage(): Did not recognize command"
	}
	log.Printf(msg)
	return message.Response{
		Timestamp: time.Now(),
		Success:   success,
		Message:   msg,
	}
}

// Start accepts a configuration KVP object, and starts both the game engine and API server
func Start(config ...map[string]string) {
	initialize()
	configure(config)
	startAPI()
}
