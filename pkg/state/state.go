package state

import "github.com/btrump/taurus-server/pkg/phase"

type State struct {
	Order        []string
	RoundCounter int
	TurnCounter  int
	Phase        phase.Phase
}
