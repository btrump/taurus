package ttt

// State describes the current game world
type State struct {
	Players      [2]*Player
	Order        [2]int
	RoundCounter int
	TurnCounter  int
	Phase        Phase
	Data         Data
	Score        [2]int
}

// Data is a container for the objects in the game world
type Data struct {
	Env [9]string
}

func (e *Engine) isTurn(id string) bool {
	return e.PlayerCurrent() == id
}

// func (f *FSM) isOpen(i string) bool {
// 	f.State.Data.Env
// 	p := f.PlayerCurrent()
// 	if p != nil {
// 		return p.ID == id
// 	}
// 	return false
// }