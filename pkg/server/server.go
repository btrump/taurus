package server

import (
	"fmt"
	"log"
	"time"

	"github.com/btrump/taurus-server/internal/helper"
	"github.com/btrump/taurus-server/internal/message"
	"github.com/btrump/taurus-server/pkg/client"
	"github.com/btrump/taurus-server/pkg/phase"
	"github.com/btrump/taurus-server/pkg/state"
	"github.com/google/uuid"
)

// Config is a container for transient server settings
type Config struct {
}

// Connection holds information about a connected client
type Connection struct {
	ID        string
	Connected time.Time
	Client    *client.Client
}

// Server is an instance of the Taurus server
type Server struct {
	Name     string
	Version  string
	ID       string
	Config   Config
	Started  time.Time
	Clients  []Connection
	Messages []interface{}
	Chat     []string
	State    state.State
}

// initialize sets the initial, static server values
func (s *Server) initialize() {
	s.ID = uuid.New().String()
	log.Printf("server::initialize(): Initializing new server %s", s.ID)
	s.State = state.State{
		Phase: phase.PRE,
	}
}

// configure sets the transient server values
func (s *Server) configure(config []map[string]string) {
	s.Started = time.Now()
	s.Name = "taurus-server"
	s.Version = "development"
	log.Printf("server::configureServer(): %s", helper.ToJSON(s.Config))
}

// Status returns the current status of the server and the engine state
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
		Clients      []Connection
		Messages     []interface{}
		State        state.State
	}{s.ID, s.Name, s.Version, s.Started, time.Now().Sub(s.Started), len(s.Clients), len(s.Messages), len(s.Chat), s.State.TurnCounter, s.State.RoundCounter, s.State.Phase, s.Config, s.Clients, s.Messages, s.State}
	return helper.ToJSON(status)
}

// ClientConnect adds a client to the list of clients and to the order list
func (s *Server) ClientConnect(client client.Client) (message.Response, error) {
	log.Printf("server::ClientConnect(): %s connected", client.ID)
	log.Printf("server::ClientConnect(): appending '%s' to client list", client.ID)
	s.Clients = append(s.Clients, Connection{
		Client:    &client,
		Connected: time.Now(),
	})
	log.Printf("server::ClientConnect(): appending '%s' to order list", client.ID)
	s.State.Order = append(s.State.Order, client.ID)
	return message.NewResponse(true, fmt.Sprintf("server::ClientConnect(): %s successfully connected", client.ID)), nil
}

// requestEvaluate determines if a request is valid and, if so, handles it
func (s *Server) requestEvaluate(m message.Request) message.Response {
	if err := s.requestValidate(m); err != nil {
		log.Printf("server::evaluateMessage(): requestValidate failure")
		return message.Response{
			Timestamp: time.Now(),
			Success:   false,
			Message:   err.Error(),
		}
	}
	return s.requestExecute(m)
}

// New accepts a configuration KVP object, and returns a new configured server
func New(config ...map[string]string) Server {
	s := Server{}
	s.initialize()
	s.configure(config)
	return s
}
