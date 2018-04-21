package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bitDecayGames/LudumDare41/server/pubsub"
	"github.com/gorilla/mux"
)

const (
	port = 8080
)

var pubSubService pubsub.PubSubService

func main() {
	host := fmt.Sprintf(":%v", port)
	log.Printf("Starting server on %s ...", host)

	pubSubService = pubsub.NewPubSubService()

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/ping", PingHandler).Methods("POST")
	r.HandleFunc("/api/v1/pubsub", PubSubHandler)
	r.HandleFunc("/api/v1/pubsub/connection/{connectionID}", UpdatePubSubConnectionHandler).Methods("PUT")

	log.Printf("Server started on %s", host)

	http.ListenAndServe(host, r)
}

type PingMessage struct {
	Status string `json:"status"`
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	msg := pubsub.Message{
		Type: "ping",
	}

	errors := pubSubService.SendMessage("test", msg)
	if len(errors) > 0 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	pingMsg := PingMessage{
		Status: "ok",
	}

	pingBytes, err := json.Marshal(pingMsg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(pingBytes)

	log.Println("Ping")
}

func PubSubHandler(w http.ResponseWriter, r *http.Request) {
	connectionID, err := pubSubService.AddSubscription(w, r)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Added new pubsub subscription with connectionID %s", connectionID)
}

type UpdateSubBody struct {
	gameName   string
	playerName string
}

func UpdatePubSubConnectionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var body UpdateSubBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	connectionID := vars["connectionID"]

	err = pubSubService.UpdateSubscription(connectionID, body.gameName, body.playerName)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
