package api

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func add(a int, b int) int {
	return a + b
}

func mod(a int, b int) int {
	return a % b
}

func col(i int) int {
	return i % 3
}

func row(i int) int {
	return int(i / 3)
}

func (a *API) score(i int) int {
	return 0 //a.Server.Engine.GetScore(i)
	// State.Data.Score[i]
}

func (a *API) isConnected(i int) bool {
	// return 0 //a.Server.Engine.GetScore(i)
	// State.Data.Score[i]
	// return a.Server.Engine.GetState().Players[i]
	return true
}

func (a *API) currentPlayer() string {
	p := a.Server.Engine.PlayerCurrent()
	if p == "" {
		return "No players connected"
	}
	return fmt.Sprintf("%s's turn", p)
}

func (a *API) ttt(w http.ResponseWriter, r *http.Request) {
	// tFile := "pkg/api/templates/tic-tac-toe.html"
	tFile := "tic-tac-toe.html"
	funcMap := template.FuncMap{
		"mod":           mod,
		"add":           add,
		"row":           row,
		"col":           col,
		"currentPlayer": a.currentPlayer,
		"score":         a.score,
		"isConnected":   a.isConnected,
	}
	t, _ := template.New(tFile).Funcs(funcMap).ParseFiles(tFile)
	log.Printf("%v", a.Server.Engine.GetState())
	t.Execute(w, a.Server.Engine.GetState())
}
