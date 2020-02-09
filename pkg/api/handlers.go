package api

import (
	"fmt"
	"html/template"
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

func (a *API) currentPlayer() string {
	p := a.Server.FSM.PlayerCurrent()
	if p == nil {
		return "No players connected"
	}
	return fmt.Sprintf("%s's turn", p.Name)
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
	}
	t, _ := template.New(tFile).Funcs(funcMap).ParseFiles(tFile)
	t.Execute(w, a.Server.FSM.State)
}
