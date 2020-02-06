package message

import (
	"time"

	"github.com/google/uuid"
)

type Request struct {
	// Execute()
	// Validate()
	ID        string
	Timestamp time.Time
	UserID    string
	Command   string
	Message   string
}

type Response struct {
	ID        string
	Timestamp time.Time
	Success   bool
	Message   string
}

func NewResponse(s bool, m string) Response {
	return Response{
		ID:        uuid.New().String(),
		Timestamp: time.Now(),
		Success:   s,
		Message:   m,
	}
}
