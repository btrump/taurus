package fsm

// State describes the current game world
type State struct {
	Players      map[string]*Player
	Order        []string
	RoundCounter int
	TurnCounter  int
	Phase        Phase
	Data         Data
}

// Data is a container for the objects in the game world
type Data struct {
	Env   [9]string
	Score [2]int
}

func (f *FSM) isTurn(id string) bool {
	p := f.PlayerCurrent()
	if p != nil {
		return p.ID == id
	}
	return false
}

// func (f *FSM) isOpen(i string) bool {
// 	f.State.Data.Env
// 	p := f.PlayerCurrent()
// 	if p != nil {
// 		return p.ID == id
// 	}
// 	return false
// }
