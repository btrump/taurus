/*
Package api provides an implementation of an http API that faciliates
communication between the server and taurus clients
*/
package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/btrump/taurus-server/internal/helper"
	"github.com/btrump/taurus-server/pkg/client"
	"github.com/btrump/taurus-server/pkg/message"
	"github.com/btrump/taurus-server/pkg/server"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Clients is a dev fixture for testing with clients
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

// API is an instance of the Taurus api, facilitates communication between client and server
type API struct {
	ID            string
	Router        *mux.Router
	Server        *server.Server
	Version       string
	BytesSent     int
	BytesReceived int
	Port          int
}

// New returns a new API
func New() API {
	a := API{
		ID:      uuid.New().String(),
		Port:    8081,
		Version: "development",
	}
	log.Printf("api::New(): New API %s", helper.ToJSON(a))
	return a
}

// attachRouter associates a router with the API
func (a *API) attachRouter() {
	a.Router = mux.NewRouter().StrictSlash(true)
	a.Router.HandleFunc("/api", a.requestParse).Methods("POST")
	a.Router.HandleFunc("/status/api", a.status)
	a.Router.HandleFunc("/status/server", a.serverStatus)
	a.Router.HandleFunc("/client", a.getClients)
	a.Router.HandleFunc("/client/{id}", a.getClient)
	a.Router.HandleFunc("/client/{id}/connect", a.clientConnect)
	a.Router.HandleFunc("/ttt", a.ttt)
}

// Use associates a server with the API
func (a *API) Use(s *server.Server) {
	log.Printf("api::Use(): Using server %s", s.ID)
	a.Server = s
	log.Printf("api::Use(): Attaching router")
	a.attachRouter()
}

// clientConnect connects a client to a server
func (a *API) clientConnect(w http.ResponseWriter, r *http.Request) {
	payload, _ := json.Marshal(a)
	log.Printf("API object: %s", payload)
	vars := mux.Vars(r)
	m, err := a.Server.ClientConnect(Clients[vars["id"]])
	if err != nil {
		a.sendJSON(err, w)
	}
	a.sendJSON(m, w)
}

// getClient sends the requested client
func (a *API) getClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	a.sendJSON(Clients[vars["id"]], w)
}

// getClients sends the list of currently connected clients
func (a *API) getClients(w http.ResponseWriter, r *http.Request) {
	a.sendJSON(Clients, w)
}

// requestExecute sends validated requests to the associated server
func (a *API) requestExecute(m message.Request) message.Response {
	return a.Server.ProcessRequest(m)
}

// requestParse determines if a request is valid and, if so, handles it
func (a *API) requestParse(w http.ResponseWriter, r *http.Request) {
	var m message.Request
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	m.Timestamp = time.Now()
	var response = a.requestExecute(m)
	log.Printf("api::requestParse(): Got request %s", m)
	a.sendJSON(response, w)
}

// serverStatus returns status information about the associated server
func (a *API) serverStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(a.Server.Status()))
}

// sendJSON serializes an object as JSON and writes it to http
func (a *API) sendJSON(v interface{}, w http.ResponseWriter) int {
	payload, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
	log.Printf("server::sendJSON(): SEND %d bytes: %s", len(payload), payload)
	a.BytesSent += len(payload)
	return len(payload)
}

// Start begins serving the API on the configured port
func (a *API) Start() {
	log.Printf("api::Start(): Listening on 0.0.0.0:%d", a.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", a.Port), a.Router))
}

// status returns information about the API
func (a *API) status(w http.ResponseWriter, r *http.Request) {
	a.sendJSON(struct {
		ID            string
		Version       string
		Port          int
		BytesSent     int
		BytesReceived int
	}{a.ID, a.Version, a.Port, a.BytesSent, a.BytesReceived}, w)
}
