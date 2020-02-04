package server

import (
	"log"
	"time"

	"github.com/btrump/taurus-server/internal/helper"
	"github.com/btrump/taurus-server/internal/message"
	"github.com/btrump/taurus-server/pkg/client"
	"github.com/btrump/taurus-server/pkg/phase"
	"github.com/btrump/taurus-server/pkg/state"
	"github.com/google/uuid"
)

type Config struct {
}

type clientConnection struct {
	ID        string
	Connected time.Time
	Client    *client.Client
}

type Server struct {
	Name     string
	Version  string
	ID       string
	Config   Config
	Started  time.Time
	Clients  []clientConnection
	Messages []interface{}
	Chat     []string
	State    state.State
}

func (s *Server) initialize() {
	s.ID = uuid.New().String()
	log.Printf("server::initialize(): Initializing new server %s", s.ID)
	s.State = state.State{
		Phase: phase.PHASE_PRE,
	}
}

func (s *Server) configure(config []map[string]string) {
	s.Started = time.Now()
	s.Name = "taurus-server"
	s.Version = "development"
	log.Printf("server::configureServer(): %s", helper.ToJSON(s.Config))
}

func (s *Server) Status() string {
	status := struct {
		ID           string
		Name         string
		Version      string
		Started      time.Time
		Uptime       time.Duration
		ClientCount  int
		MessageCount int
		ChatCount    int
		TurnCounter  int
		RoundCounter int
		Phase        phase.Phase
		Config       Config
		Clients      []clientConnection
		Messages     []interface{}
		State        state.State
	}{s.Name, s.Version, s.ID, s.Started, time.Now().Sub(s.Started), len(s.Clients), len(s.Messages), len(s.Chat), s.State.TurnCounter, s.State.RoundCounter, s.State.Phase, s.Config, s.Clients, s.Messages, s.State}
	return helper.ToJSON(status)
}

func (s *Server) ClientConnect(client client.Client) {
	log.Printf("server::ClientConnect(): %s connected", client.ID)
	log.Printf("server::ClientConnect(): appending '%s' to client list", client.ID)
	s.Clients = append(s.Clients, clientConnection{
		Client:    &client,
		Connected: time.Now(),
	})
	log.Printf("server::ClientConnect(): appending '%s' to order list", client.ID)
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
	s.initialize()
	s.configure(config)
	return s
}
