package phase

// Phase is the current phase of the game state
type Phase int

// Valid phases
const (
	PRE Phase = iota
	STARTED
	COMPLETED
)
