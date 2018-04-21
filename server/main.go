package main

import (
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

	log.Printf("Server started on %s", host)

	http.ListenAndServe(host, r)
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	msg := pubsub.Message{
		Type: "ping",
	}

	err := pubSubService.SendMessage(msg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	log.Println("Ping")
}

func PubSubHandler(w http.ResponseWriter, r *http.Request) {
	err := pubSubService.AddSubscription(w, r)
	if err != nil {
		log.Println(err)
	}
	log.Println("Added pubsub subscription")
}
