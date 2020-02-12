/*
Package message implements a basic envelope for passing messages from client
to API to server to FSM
*/
package message

import (
	"time"

	"github.com/google/uuid"
)

// Request is a container for a client-to-server message
type Request struct {
	ID        string
	Timestamp time.Time
	UserID    string
	Command   string
	Message   string
}

// Response is a container for a server-to-client message
type Response struct {
	ID        string
	Timestamp time.Time
	Success   bool
	Message   string
}

// NewResponse accepts a success code and string, and returns a stamped response
func NewResponse(s bool, m string) Response {
	return Response{
		ID:        uuid.New().String(),
		Timestamp: time.Now(),
		Success:   s,
		Message:   m,
	}
}

// NewRequest accepts a success code and string, and returns a stamped response
func NewRequest(u string, c string, m string) Request {
	return Request{
		ID:        uuid.New().String(),
		Timestamp: time.Now(),
		UserID:    u,
		Command:   c,
		Message:   m,
	}
}
