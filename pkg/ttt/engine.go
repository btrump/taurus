package ttt

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/btrump/taurus-server/pkg/message"
)

// Phase is the current phase of the game state
type Phase int

// Valid phases
const (
	PRE Phase = iota
	STARTED
	COMPLETED
)

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

// Engine is a container for the state of the game world
type Engine struct {
	state State
}

// New returns a new, initialized state machine
func New() *Engine {
	log.Printf("fsm::New(): Creating new Engine")
	e := &Engine{}
	e.initialize()
	return e
}

// PlayerCurrent returns the currently active player
func (e Engine) PlayerCurrent() string {
	return e.state.Players[e.state.Order[e.state.TurnCounter%len(e.state.Order)]].ID
}

// PlayerAdd adds a player to the list of players and the order list
func (e *Engine) PlayerAdd(n string) (message.Response, error) {
	l := len(e.GetPlayers())
	if l > 1 {
		msg := "ttt::PlayerAdd(): Failed to add player %s"
		return message.NewResponse(false, fmt.Sprintf(msg, n)), errors.New(msg)
	}
	e.state.Players = append(e.state.Players, NewPlayer(n, n))
	e.state.Order = append(e.state.Order, l)
	return message.NewResponse(true, fmt.Sprintf("ttt::PlayerAdd(): Added player %s", n)), nil
}

// GetPlayers returns all connected players
func (e *Engine) GetPlayers() []*Player {
	return e.state.Players
}

// GetState returns the current engine state
func (e *Engine) GetState() interface{} {
	return e.state
}

// Stats returns statistics about the engine
func (e *Engine) Stats() interface{} {
	return struct{}{}
}

// SetScore sets the score for a given player
func (e *Engine) SetScore(p int, s int) int {
	e.state.Score[p] = s
	return s
}

// GetScore returns the score for player i
func (e *Engine) GetScore(i int) int {
	return e.state.Score[i]
}

// SetPhase sets the engine to a given phase
func (e *Engine) SetPhase(p Phase) {
	e.state.Phase = p
}

// initialize sets the initial values of the state machine
func (e *Engine) initialize() {
	log.Printf("fsm::initialize(): Initializing Engine")
	e.state = State{
		Phase: PRE,
	}
	e.state.Players = make([]*Player, 0, 2)
	e.state.Order = make([]int, 0, 2)
}

// Validate ensures that a request is valid
func (e *Engine) Validate(m message.Request) (message.Response, error) {
	var err error
	switch m.Command {
	case "GAME_START":
		if e.state.Phase != PRE {
			err = errors.New("server::Validate(): Could not start game")
		}
	case "GAME_END":
		if e.state.Phase == COMPLETED {
			err = errors.New("server::Validate(): Could not end game. Game already ended")
		}
	case "TURN_END":
		if false {
			err = errors.New("server::Validate(): Could not end turn")
		}
	case "NEXT_PHASE":
		if e.state.Phase != STARTED {
			err = errors.New("server::Validate(): Could not advance phase. Not in STARTED state")
		}
	case "MARK_TILE":
	default:
		err = errors.New("server::Validate(): Did not recognize command")
	}
	if err != nil {
		return message.NewResponse(false, err.Error()), err
	}
	return message.NewResponse(true, "server::Validate(): Valid command"), err
}

// Execute performs the command indicated by a request
func (e *Engine) Execute(m message.Request) (message.Response, error) {
	var err error
	var msg string
	switch m.Command {
	case "GAME_START":
		e.state.Phase = STARTED
		msg = "server::requestExecute(): Game started"
	case "GAME_END":
		e.state.Phase = COMPLETED
		msg = "server::requestExecute(): Game ended"
	case "TURN_END":
		if !e.IsTurn(m.UserID) {
			msg = "server::requestExecute(): Not player's turn"
			err = errors.New(msg)
			break
		}
		e.state.TurnCounter++
		msg = "server::requestExecute(): Ending turn"
		if e.state.TurnCounter%len(e.state.Order) == 0 {
			msg += ". Ending round."
			e.state.RoundCounter++
		}
	case "NEXT_PHASE":
		e.state.Phase++
		msg = "server::requestExecute(): Advancing to next phase"
	case "MARK_TILE":
		tile, _ := strconv.ParseInt(m.Message, 0, 64)
		e.state.Data.Env[tile] = m.UserID
	default:
		msg = "server::requestExecute(): Did not recognize command. This should never happen, because request was already validated"
	}
	return message.NewResponse(err == nil, msg), err
}

// IsTurn returns true if it is the supplied player's turn, false otherwise
func (e *Engine) IsTurn(id string) bool {
	return e.PlayerCurrent() == id
}

// GetPhase returns the engine's current phase
func (e *Engine) GetPhase() Phase {
	return e.state.Phase
}
