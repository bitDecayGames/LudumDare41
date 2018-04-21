package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/bitDecayGames/LudumDare41/server/pubsub"
	"github.com/gorilla/mux"
)

const (
	port       = 8080
	apiv1      = "/api/v1"
	lobbyRoute = apiv1 + "/lobby"
)

var pubSubService pubsub.PubSubService

func main() {
	host := fmt.Sprintf(":%v", port)
	log.Printf("Starting server on %s ...", host)

	rand.Seed(time.Now().UnixNano())

	pubSubService = pubsub.NewPubSubService()

	r := mux.NewRouter()
	// Test ping
	r.HandleFunc(apiv1+"/ping", PingHandler).Methods("POST")
	// PubSub
	r.HandleFunc(apiv1+"/pubsub", PubSubHandler)
	r.HandleFunc(apiv1+"/pubsub/connection/{connectionID}", UpdatePubSubConnectionHandler).Methods("PUT")
	// Lobby
	// r.HandleFunc(lobbyRoute, LobbyCreateHandler).Methods("POST")
	// r.HandleFunc(lobbyRoute+"{lobbyName}/join", LobbyJoinHandler).Methods("POST")
	// r.HandleFunc(lobbyRoute+"{lobbyName}/start", LobbyStartHandler).Methods("PUT")

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

type updateSubBody struct {
	GameName   string `json:"gameName"`
	PlayerName string `json:"playerName"`
}

func UpdatePubSubConnectionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var body updateSubBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	connectionID := vars["connectionID"]

	err = pubSubService.UpdateSubscription(connectionID, body.GameName, body.PlayerName)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

// func LobbyCreateHandler(w http.ResponseWriter, r *http.Request) {

// }

// type joinLobbyBody struct {
// 	PlayerName string `json:"playerName"`
// }

// func LobbyJoinHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// }

// func LobbyStartHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// }
