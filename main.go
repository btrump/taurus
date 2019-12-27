package main

import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"

  "github.com/gorilla/mux"
)

type Client struct {
  ID      string `json:"id"`
  Name    string `json:"name"`
}

type Message struct {
  ID string `json:"id"`
  Message string `json:"message"`
}

var clients = map[string]Client{
  "1": {
      ID: "1",
      Name: "Blair Trump",
    },
  "2": {
      ID: "2",
      Name: "Mina Hu",
    },
  }

func about(w http.ResponseWriter, r *http.Request) {
  apiVersion := "v0.1"
  fmt.Fprintf(w, "API %s", apiVersion)
}

func getClient(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  payload, _ := json.Marshal(clients[vars["id"]])
  w.Header().Set("Content-Type", "application/json")
  w.Write(payload)
  fmt.Println(clients)
}

func getClients(w http.ResponseWriter, r *http.Request) {
  payload, _ := json.Marshal(clients)
  w.Header().Set("Content-Type", "application/json")
  w.Write(payload)
  fmt.Println(clients)
}

func parseAction(w http.ResponseWriter, r *http.Request) {
  var m Message
  err := json.NewDecoder(r.Body).Decode(&m)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }
  fmt.Fprintf(w, "Message: %+v", m)
  fmt.Println("Message: %+v", m)
}

func main() {
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/action", parseAction)
  router.HandleFunc("/about", about)
  router.HandleFunc("/client", getClients)
  router.HandleFunc("/client/{id}", getClient)
  port := 8081
  msg := fmt.Sprintf("Listening on port %d", port)
  fmt.Println(msg)
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
