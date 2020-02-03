package message

import "time"

type Request struct {
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
