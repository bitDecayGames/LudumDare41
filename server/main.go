package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"time"

	"github.com/bitDecayGames/LudumDare41/server/lobby"

	"github.com/bitDecayGames/LudumDare41/server/pubsub"
	"github.com/gorilla/mux"
)

const (
	port       = 8080
	apiv1      = "/api/v1"
	lobbyRoute = apiv1 + "/lobby"
	gameRoute  = apiv1 + "/game"
)

var pubSubService pubsub.PubSubService
var lobbyService lobby.LobbyService

func main() {
	host := fmt.Sprintf(":%v", port)
	log.Printf("Starting server on %s ...", host)

	rand.Seed(time.Now().UnixNano())

	pubSubService = pubsub.NewPubSubService()
	lobbyService = lobby.NewLobbyService()

	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	// Test ping
	r.HandleFunc(apiv1+"/ping", PingHandler).Methods("POST")
	// PubSub
	r.HandleFunc(apiv1+"/pubsub", PubSubHandler)
	r.HandleFunc(apiv1+"/pubsub/connection/{connectionID}", UpdatePubSubConnectionHandler).Methods("PUT")
	// Lobby
	r.HandleFunc(lobbyRoute, LobbyCreateHandler).Methods("POST")
	r.HandleFunc(lobbyRoute+"/{lobbyName}/join", LobbyJoinHandler).Methods("POST")
	// TODO Below
	r.HandleFunc(lobbyRoute+"/{lobbyName}/players", LobbyJoinHandler).Methods("GET")
	// Don't return anything
	r.HandleFunc(lobbyRoute+"/{lobbyName}/start", LobbyStartHandler).Methods("PUT")
	// Game
	// Cards are list on ints, need tick as well
	// Trigger next round once all submitted
	r.HandleFunc(lobbyRoute+"/{gameName}/tick/{tick}/player/{playerName}/cards", CardsSubmitHandler).Methods("PUT")
	// Get current tick
	r.HandleFunc(lobbyRoute+"/{gameName}/tick", GetCurrentTickHandler).Methods("GET")
	// Game state + players cards for a tick
	r.HandleFunc(lobbyRoute+"/{gameName}/tick/{tick}/player/{playerName}", GetGameStateHandler).Methods("GET")

	log.Printf("Server started on %s", host)

	http.ListenAndServe(host, r)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Save a copy of this request for debugging.
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Println(err)
			next.ServeHTTP(w, r)
			return
		}

		rec := httptest.NewRecorder()
		next.ServeHTTP(rec, r)

		responseDump, err := httputil.DumpResponse(rec.Result(), true)
		if err != nil {
			log.Println(err)
			return
		}

		// we copy the captured response headers to our new response
		for k, v := range rec.Header() {
			w.Header()[k] = v
		}

		// grab the captured response body
		data := rec.Body.Bytes()
		w.Write(data)

		log.Printf("%s\n\nRESPONSE\n%s", requestDump, responseDump)
	})
}

type pingBody struct {
	GameName string `json:"gameName"`
}

type PingMessage struct {
	Status string `json:"status"`
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	var body pingBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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
	connectionID := vars["connectionID"]

	var body updateSubBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = pubSubService.UpdateSubscription(connectionID, body.GameName, body.PlayerName)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

type newLobbyResBody struct {
	Name string `json:"name"`
}

func LobbyCreateHandler(w http.ResponseWriter, r *http.Request) {
	lobby, err := lobbyService.NewLobby()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resBody := newLobbyResBody{
		Name: lobby.Name,
	}
	err = json.NewEncoder(w).Encode(&resBody)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type joinLobbyReqBody struct {
	PlayerName string `json:"playerName"`
}

type joinLobbyResBody struct {
	SanitizedPlayerName string `json:"sanitizedPlayerName"`
}

// Require a valid name, otherwise reject. return santizied version
func LobbyJoinHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lobbyName := vars["lobbyName"]

	var reqBody joinLobbyReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	lobby, err := lobbyService.GetLobby(lobbyName)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	sanitizedPlayerName, err := lobby.AddPlayer(reqBody.PlayerName)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resBody := joinLobbyResBody{
		SanitizedPlayerName: sanitizedPlayerName,
	}
	err = json.NewEncoder(w).Encode(&resBody)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func LobbyGetPlayersHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
}

func LobbyStartHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
}

func CardsSubmitHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
}

func GetCurrentTickHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
}

func GetGameStateHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
}
