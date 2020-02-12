package ttt_test

import (
	"testing"

	"github.com/btrump/taurus-server/pkg/ttt"
)

/*
TestNewPlayer expects a newly-initiatlized player
*/
func TestNewPlayer(t *testing.T) {
	var n string = "test_name"
	var id string = "test_id"
	p := ttt.NewPlayer(id, n)
	v := ttt.Player{
		ID:   id,
		Name: n,
	}
	if *p != v {
		t.Errorf("Got '%v'; Want '%v'", p, v)
	}
}
