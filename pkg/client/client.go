/*
Package client implements a client which communicates with the taurus server
*/
package client

import (
	"bufio"
	"net"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	// logrus.SetFormatter(&logrus.JSONFormatter{})
	log = logrus.WithFields(
		logrus.Fields{
			"application": "taurus-client",
			"version":     "0.0.1",
		},
	)
}

// Client is a container holding client information
type Client struct {
	id   string
	conn net.TCPConn
}

// New returns a new instance of a client
func New() *Client {
	c := Client{}
	c.initialize()
	log.Printf("Created client %v", c.id)
	return &c
}

func (c *Client) initialize() {
	c.id = uuid.New().String()
}

// Connect connects to socket
func (c *Client) Connect(socket string) {
	log.Printf("Connecting to %v", socket)
	conn, err := net.Dial("tcp", socket)
	if err != nil {
		log.Println(err)
		return
	}
	message, _ := bufio.NewReader(conn).ReadString('\n')
	log.Printf("Got msg: '%v'", message)
}
