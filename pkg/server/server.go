package server

import (
	"fmt"
	"log"
	"time"

	"github.com/btrump/taurus/internal/message"
	"github.com/btrump/taurus/pkg/client"
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
}

var Config serverConfig
var Clients []clientConnection
var Messages []interface{}
var Chat []string
var State state

func configureServer(config []map[string]string) serverConfig {
	log.Printf("server::configureServer(): PLACEHOLDER")
	return serverConfig{
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

func receiveRequest(m message.Request) string {
	var res = fmt.Sprintf("server::receiveRequest(): Got message with command '%s' from user '%s'", m.Command, m.UserID)
	log.Print(res)
	Messages = append(Messages, m)
	return res
}

// func evaluateMessage(message message.Request) {
//
// }

// Start accepts a configuration KVP object, and starts both the game engine and API server
func Start(config ...map[string]string) {
	Config = configureServer(config)
	startAPI()
}
