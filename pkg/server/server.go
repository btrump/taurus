package server

import (
	"log"
	"time"

	"github.com/btrump/taurus-server/internal/message"
	"github.com/btrump/taurus-server/pkg/client"
	"github.com/btrump/taurus-server/pkg/phase"
	"github.com/btrump/taurus-server/pkg/state"
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

func (s *Server) evaluateMessage(m message.Request) message.Response {
	if err := s.requestValidate(m); err != nil {
		log.Printf("server::evaluateMessage(): requestValidate failure")
		return message.Response{
			Timestamp: time.Now(),
			Success:   false,
			Message:   err.Error(),
		}
	}
	return s.requestHandle(m)
}

// New accepts a configuration KVP object, and starts both the game engine and API server
func New(config ...map[string]string) Server {
	s := Server{}
	initialize(&s)
	configure(&s, config)
	return s
}
