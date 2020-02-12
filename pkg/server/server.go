/*
Package server providdes an implementation of a server that accepts client
requests and passes them to an underlying state machine for evaluation
*/
package server

import (
	"fmt"
	"log"
	"time"

	"github.com/btrump/taurus-server/internal/helper"
	"github.com/btrump/taurus-server/pkg/client"
	"github.com/btrump/taurus-server/pkg/engine"
	"github.com/btrump/taurus-server/pkg/message"
	"github.com/btrump/taurus-server/pkg/ttt"
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
	Engine   engine.Engine
}

// initialize sets the initial, static server values
func (s *Server) initialize() {
	s.ID = uuid.New().String()
	log.Printf("server::initialize(): Initializing new server %s", s.ID)
	s.Engine = ttt.New()
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
		ID            string
		Name          string
		Version       string
		Started       time.Time
		Uptime        time.Duration
		ClientCount   int
		MessageCount  int
		ChatCount     int
		CurrentPlayer string
		Config        Config
		Clients       []Connection
		Messages      []interface{}
		Stats         interface{}
	}{s.ID, s.Name, s.Version, s.Started, time.Now().Sub(s.Started), len(s.Clients), len(s.Messages), len(s.Chat), s.Engine.PlayerCurrent(), s.Config, s.Clients, s.Messages, s.Engine.Stats()}
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
	s.Engine.PlayerAdd(client.Name)
	return message.NewResponse(true, fmt.Sprintf("server::ClientConnect(): %s successfully connected", client.ID)), nil
}

// ProcessRequest accepts requests, stamps them with IDs
func (s *Server) ProcessRequest(req message.Request) message.Response {
	req.ID = uuid.New().String()
	log.Printf("server::ProcessRequest(): Got message with command '%s' from user '%s'. Assigned id %s", req.Command, req.UserID, req.ID)
	res, err := s.Engine.Validate(req)
	if err != nil {
		log.Printf("server::ProcessRequest(): request %s is not valid", res.ID)
	} else {
		log.Printf("server::ProcessRequest(): request %s is valid", res.ID)
		res, _ = s.Engine.Execute(req)
	}
	s.Messages = append(s.Messages, struct {
		Request  message.Request
		Response message.Response
	}{req, res})
	return res
}

// New accepts a configuration KVP object, and returns a new configured server
func New(config ...map[string]string) Server {
	s := Server{}
	s.initialize()
	s.configure(config)
	return s
}
