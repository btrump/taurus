package message_test

import (
	"testing"

	"github.com/btrump/taurus-server/pkg/message"
)

func TestNewResponse(t *testing.T) {
	msg := "test message"
	s := true
	r1 := message.Response{
		Success: s,
		Message: msg,
	}
	var r2 = message.NewResponse(s, msg)
	if (r1.Success != r2.Success) || (r1.Message != r2.Message) {
		t.Errorf("Got '%v' and %s'; want '%v' and '%s'", r1.Success, r1.Message, r2.Success, r2.Message)
	}
}

func TestNewRequest(t *testing.T) {
	msg := "test message"
	id := "1"
	cmd := "test command"
	r1 := message.Request{
		UserID:  id,
		Command: cmd,
		Message: msg,
	}
	var r2 = message.NewRequest(id, cmd, msg)
	if (r1.UserID != r2.UserID) || (r1.Command != r2.Command) || (r1.Message != r2.Message) {
		t.Errorf("Got '%s', '%s', '%s'; want '%s', '%s', '%s'", r1.UserID, r1.Command, r1.Message, r2.UserID, r2.Command, r2.Message)
	}
}
