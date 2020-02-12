package ttt

// State describes the current game world
type State struct {
	Players      []*Player
	Order        []int
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

func (e *Engine) IsTurn(id string) bool {
	return e.PlayerCurrent() == id
}
