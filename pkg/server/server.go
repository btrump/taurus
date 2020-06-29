package server

import (
	"fmt"
	"net"
	"time"

	"github.com/google/uuid"

	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(
		logrus.Fields{
			"application": "taurus-server",
			"version":     "0.0.1",
		},
	)
}

// Server is an instance of the Taurus server
type Server struct {
	count    int
	listener net.Listener
	port     int
	clients  [10]int
	started  time.Time
	id       string
}

// New returns a new instance of a server
func New() *Server {
	s := Server{}
	s.initialize()
	return &s
}

func (s *Server) initialize() {
	s.count = 0
	s.port = 8080
	s.started = time.Now()
	socket := fmt.Sprintf("localhost:%d", s.port)
	s.listener, _ = net.Listen("tcp", socket)
	s.id = uuid.New().String()
}

func (s *Server) handleConnection(conn net.Conn) {
	s.count = s.count + 1
	log.Printf("Handling req #%v", s.count)
	log.Printf("Connection from remote with addr: %v", conn.RemoteAddr())
	msg := time.Now().String()
	conn.Write([]byte(msg))
	conn.Close()
}

func (s *Server) handleError(err error) {
	log.Printf("There was an error: %v", err)
}

func (s *Server) listen() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			s.handleError(err)
		}
		go s.handleConnection(conn)
	}
}

// Start makes the server begin listening on port s.port
func (s *Server) Start() {
	log.Printf("Server started at %v", s.started)
	log.Printf("Listening on %d", s.port)
	s.listen()
}
