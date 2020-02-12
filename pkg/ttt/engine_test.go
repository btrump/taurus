package ttt_test

import (
	"testing"

	"github.com/btrump/taurus-server/pkg/message"
	"github.com/btrump/taurus-server/pkg/ttt"
)

/*
TestNew expects a newly-initiatlized Engine that:
- Has 0 players
- Has an empty, 9-tile environment
- Has a zero score
- is in the PRE phase
*/
func TestNew(t *testing.T) {
	e := ttt.New()
	p := len(e.GetPlayers())
	s := e.GetState().(ttt.State)
	if p != 0 {
		t.Errorf("Got '%d' players; want 0", p)
	}
	if len(s.Data.Env) != 9 {
		t.Errorf("Got '%d' tiles; want 9", len(s.Data.Env))
	}
	for i, v := range s.Data.Env {
		if v != "" {
			t.Errorf("Got value %s for tile %d; want 0", v, i)
		}
	}
	for i, v := range s.Score {
		if v != 0 {
			t.Errorf("Got score %v for player %v; wanted 0", v, i)
		}
	}
	if s.Phase != ttt.PRE {
		t.Errorf("Got phase %v; want %v", s.Phase, ttt.PRE)
	}
}

/*
TestIsTurn adds two players (a and b). It then expects the current player to be
A, ends A's turn, then expects the current player to be B. Ends B's turn,
then expects current player to be A again.
*/
func TestIsTurn(t *testing.T) {
	e := ttt.New()
	e.PlayerAdd("a")
	e.PlayerAdd("b")
	m := message.NewRequest("a", "TURN_END", "")
	if e.PlayerCurrent() != "a" {
		t.Errorf("Got player '%v'; want 'a'", e.PlayerCurrent())
	}

	e.Execute(m)
	if e.PlayerCurrent() != "b" {
		t.Errorf("Got player '%v'; want 'b'", e.PlayerCurrent())
	}

	m = message.NewRequest("b", "TURN_END", "")
	e.Execute(m)
	if e.PlayerCurrent() != "a" {
		t.Errorf("Got player '%v'; want 'a'", e.PlayerCurrent())
	}
}

func TestGetScore(t *testing.T) {
	e := ttt.New()
	e.PlayerAdd("a")
	e.PlayerAdd("b")
	if e.SetScore(0, 1) != e.GetScore(0) {
		t.Errorf("Got %v; wanted 1", e.GetScore(0))
	}
}

func TestSetScore(t *testing.T) {
	e := ttt.New()
	e.PlayerAdd("a")
	e.PlayerAdd("b")
	if e.SetScore(0, 1) != e.GetScore(0) {
		t.Errorf("Got %v; wanted 1", e.GetScore(0))
	}
}

func TestPlayerCurrent(t *testing.T) {
	e := ttt.New()
	e.PlayerAdd("a")
	e.PlayerAdd("b")
	m := message.NewRequest("a", "TURN_END", "")
	if e.PlayerCurrent() != "a" {
		t.Errorf("Got player '%v'; want 'a'", e.PlayerCurrent())
	}

	e.Execute(m)
	if e.PlayerCurrent() != "b" {
		t.Errorf("Got player '%v'; want 'b'", e.PlayerCurrent())
	}

	m = message.NewRequest("b", "TURN_END", "")
	e.Execute(m)
	if e.PlayerCurrent() != "a" {
		t.Errorf("Got player '%v'; want 'a'", e.PlayerCurrent())
	}
}

func TestStats(t *testing.T) {
	e := ttt.New()
	if e.Stats() == nil {
		t.Errorf("Got 0; want 1")
	}
}

func TestSetPhase(t *testing.T) {
	e := ttt.New()
	e.SetPhase(ttt.STARTED)
	if e.GetPhase() != ttt.STARTED {
		t.Errorf("Got '%v'; want '%v'", e.GetPhase(), ttt.STARTED)
	}
}

func TestPlayerAdd(t *testing.T) {
	e := ttt.New()
	l := len(e.GetPlayers())
	if l != 0 {
		t.Errorf("Got '%v'; want '%v'", l, 0)
	}
	// Add 3 players, only 2 should be added
	e.PlayerAdd("a")
	e.PlayerAdd("b")
	e.PlayerAdd("c")
	l = len(e.GetPlayers())
	if l != 2 {
		t.Errorf("Got '%v'; want '%v'", l, 2)
	}
}
