package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/btrump/taurus-server/internal/helper"
	"github.com/btrump/taurus-server/internal/message"
	"github.com/btrump/taurus-server/pkg/client"
	"github.com/btrump/taurus-server/pkg/server"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var Clients = map[string]client.Client{
	"1": {
		ID:   "bosa3f4",
		Name: "client1",
	},
	"2": {
		ID:   "oitnc0d",
		Name: "client2",
	},
	"3": {
		ID:   "8eexmm0",
		Name: "client3",
	},
}

type API struct {
	ID            string
	Router        *mux.Router
	Server        *server.Server
	Version       string
	BytesSent     int
	BytesReceived int
	Port          int
}

func New() API {
	a := API{
		ID:      uuid.New().String(),
		Port:    8081,
		Version: "development",
	}
	log.Printf("api::New(): New API %s", helper.ToJSON(a))
	return a
}

func (a *API) attachRouter() {
	a.Router = mux.NewRouter().StrictSlash(true)
	a.Router.HandleFunc("/api/status", a.status)
	a.Router.HandleFunc("/server/status", a.serverStatus)
	a.Router.HandleFunc("/", a.getConfig).Methods("GET")
	a.Router.HandleFunc("/", a.parseRequest).Methods("POST")
	a.Router.HandleFunc("/client", a.getClients)
	a.Router.HandleFunc("/client/{id}", a.getClient)
	a.Router.HandleFunc("/client/{id}/connect", a.clientConnect)
}

func (a *API) Use(s *server.Server) {
	log.Printf("api::Use(): Using server %s", s.ID)
	a.Server = s
	log.Printf("api::Use(): Attaching router")
	a.attachRouter()
}

func (a *API) clientConnect(w http.ResponseWriter, r *http.Request) {
	payload, _ := json.Marshal(a)
	log.Printf("API object: %s", payload)
	vars := mux.Vars(r)
	a.Server.ClientConnect(Clients[vars["id"]])
	a.sendJSON(Clients, w)
}

func (a *API) getConfig(w http.ResponseWriter, r *http.Request) {
	a.sendJSON(a.Server.Config, w)
}

func (a *API) getClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	a.sendJSON(Clients[vars["id"]], w)
}

func (a *API) getClients(w http.ResponseWriter, r *http.Request) {
	a.sendJSON(Clients, w)
}

func (a *API) handleRequest(m message.Request) message.Response {
	return a.Server.ReceiveRequest(m)
}

func (a *API) parseRequest(w http.ResponseWriter, r *http.Request) {
	var m message.Request
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	m.Timestamp = time.Now()
	var response = a.handleRequest(m)
	log.Printf("api::parseMessage(): Got request %s", m)
	a.sendJSON(response, w)
}

func (a *API) serverStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(a.Server.Status()))
}

func (a *API) sendJSON(v interface{}, w http.ResponseWriter) int {
	payload, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
	log.Printf("server::sendJSON(): SEND %d bytes: %s", len(payload), payload)
	a.BytesSent += len(payload)
	return len(payload)
}

func (a *API) Start() {
	log.Printf("api::Start(): Listening on port %d", a.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", a.Port), a.Router))
}

func (a *API) status(w http.ResponseWriter, r *http.Request) {
	a.sendJSON(struct {
		ID            string
		Version       string
		Port          int
		BytesSent     int
		BytesReceived int
	}{a.ID, a.Version, a.Port, a.BytesSent, a.BytesReceived}, w)
}
