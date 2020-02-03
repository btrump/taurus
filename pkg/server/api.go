package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/btrump/taurus/internal/message"
	"github.com/btrump/taurus/pkg/client"
	"github.com/gorilla/mux"
)

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

func connectClientCMD(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	connectClient(client.Clients[vars["id"]])
	sendJSON(Clients, w)
	// log.Printf("server::config(): PLACEHOLDER")
}

func getConfig(w http.ResponseWriter, r *http.Request) {
	sendJSON(Config, w)
	// log.Printf("server::config(): PLACEHOLDER")
}

func getClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sendJSON(client.Clients[vars["id"]], w)
	// log.Printf("api::getClient(): PLACEHOLDER")
}

func getClients(w http.ResponseWriter, r *http.Request) {
	sendJSON(Clients, w)
	// log.Printf("api::getClients(): PLACEHOLDER")
}

func handleRequest(m message.Request) message.Response {
	var r = message.Response{
		Timestamp: time.Now(),
		Success:   true,
		Message:   receiveRequest(m),
	}
	Messages = append(Messages, r)
	return r
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
	data := struct {
		Config   serverConfig
		Clients  []clientConnection
		Messages []interface{}
		Chat     []string
		State    state
	}{Config, Clients, Messages, Chat, State}
	sendJSON(data, w)
}

func startAPI() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/status", status)
	router.HandleFunc("/server/status", serverStatus)
	router.HandleFunc("/", getConfig).Methods("GET")
	router.HandleFunc("/", parseRequest).Methods("POST")
	router.HandleFunc("/client", getClients)
	router.HandleFunc("/client/{id}", getClient)
	router.HandleFunc("/client/{id}/connect", connectClientCMD)
	log.Printf("Listening on port %d", Config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", Config.Port), router))
}

var BytesSent int
var BytesReceived int
