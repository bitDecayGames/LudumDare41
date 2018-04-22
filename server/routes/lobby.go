package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bitDecayGames/LudumDare41/server/pubsub"
	"github.com/gorilla/mux"
)

const (
	lobbyRoute = apiv1 + "/lobby"
)

type LobbyRoutes struct {
	services *Services
}

func (lr *LobbyRoutes) AddRoutes(r *mux.Router) {
	r.HandleFunc(lobbyRoute, lr.lobbyCreateHandler).Methods("POST")
	r.HandleFunc(lobbyRoute+"/{lobbyName}/join", lr.lobbyJoinHandler).Methods("PUT")
	r.HandleFunc(lobbyRoute+"/{lobbyName}/players", lr.lobbyGetPlayersHandler).Methods("GET")
	r.HandleFunc(lobbyRoute+"/{lobbyName}/start", lr.lobbyStartHandler).Methods("PUT")
}

type newLobbyResBody struct {
	Name string `json:"name"`
}

func (lr *LobbyRoutes) lobbyCreateHandler(w http.ResponseWriter, r *http.Request) {
	lobby, err := lr.services.Lobby.NewLobby()
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

func (lr *LobbyRoutes) lobbyJoinHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lobbyName := vars["lobbyName"]

	var reqBody joinLobbyReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	lobby, err := lr.services.Lobby.GetLobby(lobbyName)
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

	msg := pubsub.Message{
		MessageType: pubsub.PlayerJoinMessage,
		ID:          sanitizedPlayerName,
	}
	errors := lr.services.PubSub.SendMessage(lobbyName, msg)
	if len(errors) > 0 {
		w.WriteHeader(http.StatusInternalServerError)
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

type getPlayersResBody struct {
	Players []string `json:"players"`
}

func (lr *LobbyRoutes) lobbyGetPlayersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lobbyName := vars["lobbyName"]

	lobby, err := lr.services.Lobby.GetLobby(lobbyName)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resBody := getPlayersResBody{
		Players: lobby.GetPlayers(),
	}
	err = json.NewEncoder(w).Encode(&resBody)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (lr *LobbyRoutes) lobbyStartHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lobbyName := vars["lobbyName"]

	lobby, err := lr.services.Lobby.GetLobby(lobbyName)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	errors := lr.services.CreateGame(lobby)
	if len(errors) > 0 {
		for _, err := range errors {
			log.Println(err)
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func NewLobbyRoutes(services *Services) *LobbyRoutes {
	return &LobbyRoutes{
		services: services,
	}
}
