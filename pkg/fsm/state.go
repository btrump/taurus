package fsm

// State is a container for the objects in the game world
type State struct {
	Players      map[string]*Player
	Order        []string
	RoundCounter int
	TurnCounter  int
	Phase        Phase
}

func (f *FSM) isTurn(id string) bool {
	p := f.PlayerCurrent()
	if p != nil {
		return p.ID == id
	}
	return false
}
