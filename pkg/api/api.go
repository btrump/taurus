package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/btrump/taurus-server/internal/message"
	"github.com/btrump/taurus-server/pkg/client"
	"github.com/btrump/taurus-server/pkg/server"
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
var Router *mux.Router
var Server *server.Server
var BytesSent int
var BytesReceived int

func init() {
	Router = mux.NewRouter().StrictSlash(true)
	Router.HandleFunc("/api/status", status)
	Router.HandleFunc("/server/status", serverStatus)
	Router.HandleFunc("/", getConfig).Methods("GET")
	Router.HandleFunc("/", parseRequest).Methods("POST")
	Router.HandleFunc("/client", getClients)
	Router.HandleFunc("/client/{id}", getClient)
	Router.HandleFunc("/client/{id}/connect", clientConnect)
}

func Use(s *server.Server) {
	Server = s
}

func status(w http.ResponseWriter, r *http.Request) {
	sendJSON(BytesSent, w)
}

func sendJSON(v interface{}, w http.ResponseWriter) {
	payload, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
	log.Printf("server::sendJSON(): SEND %d bytes: %s", len(payload), payload)
	BytesSent += len(payload)
}

func clientConnect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Server.ClientConnect(Clients[vars["id"]])
	sendJSON(Clients, w)
}

func getConfig(w http.ResponseWriter, r *http.Request) {
	sendJSON(Server.Config, w)
}

func getClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sendJSON(Clients[vars["id"]], w)
}

func getClients(w http.ResponseWriter, r *http.Request) {
	sendJSON(Clients, w)
}

func handleRequest(m message.Request) message.Response {
	return Server.ReceiveRequest(m)
}

func parseRequest(w http.ResponseWriter, r *http.Request) {
	var m message.Request
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	m.Timestamp = time.Now()
	var response = handleRequest(m)
	log.Printf("api::parseMessage(): Got request %s", m)
	sendJSON(response, w)
}

func serverStatus(w http.ResponseWriter, r *http.Request) {
	sendJSON(Server, w)
}

func Start() {
	log.Printf("Listening on port %d", Server.Config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", Server.Config.Port), Router))
}
