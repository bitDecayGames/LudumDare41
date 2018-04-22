package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"strconv"
	"strings"
	"time"

	"github.com/bitDecayGames/LudumDare41/server/state"

	"github.com/bitDecayGames/LudumDare41/server/routes"

	"github.com/bitDecayGames/LudumDare41/server/pubsub"
	"github.com/gorilla/mux"
)

// TODO Turn timeout logic

const (
	// Networking
	port       = 8080
	apiv1      = "/api/v1"
	lobbyRoute = apiv1 + "/lobby"
	gameRoute  = apiv1 + "/game"

	// Game
	minNumPlayers = 2
	maxNumPlayers = 4
)

var services *routes.Services
var ritz *routes.Routes

func main() {
	host := fmt.Sprintf(":%v", port)
	log.Printf("Starting server on %s ...", host)

	rand.Seed(time.Now().UnixNano())

	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	ritz = routes.InitRoutes(r)
	services = ritz.Services

	// Lobby
	r.HandleFunc(lobbyRoute, LobbyCreateHandler).Methods("POST")
	r.HandleFunc(lobbyRoute+"/{lobbyName}/join", LobbyJoinHandler).Methods("PUT")
	r.HandleFunc(lobbyRoute+"/{lobbyName}/players", LobbyGetPlayersHandler).Methods("GET")
	r.HandleFunc(lobbyRoute+"/{lobbyName}/start", LobbyStartHandler).Methods("PUT")
	// Game
	r.HandleFunc(gameRoute+"/{gameName}/tick/{tick}/player/{playerName}/cards", CardsSubmitHandler).Methods("PUT")
	r.HandleFunc(gameRoute+"/{gameName}/tick", GetCurrentTickHandler).Methods("GET")
	r.HandleFunc(gameRoute+"/{gameName}/tick/{tick}/player/{playerName}", GetGameStateHandler).Methods("GET")

	log.Printf("Server started on %s", host)

	http.ListenAndServe(host, r)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ignore websocket route.
		reqURI := strings.Split(r.RequestURI, "?")[0]
		if reqURI == routes.PubSubRoute {
			log.Printf("Skipping logging for %s", routes.PubSubRoute)
			next.ServeHTTP(w, r)
			return
		}

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
		w.WriteHeader(rec.Code)
		for k, v := range rec.Header() {
			w.Header()[k] = v
		}

		// grab the captured response body
		data := rec.Body.Bytes()
		w.Write(data)

		log.Printf("%s\n\nRESPONSE\n%s", requestDump, responseDump)
	})
}

type newLobbyResBody struct {
	Name string `json:"name"`
}

func LobbyCreateHandler(w http.ResponseWriter, r *http.Request) {
	lobby, err := services.Lobby.NewLobby()
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

	lobby, err := services.Lobby.GetLobby(lobbyName)
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
	errors := services.PubSub.SendMessage(lobbyName, msg)
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

func LobbyGetPlayersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lobbyName := vars["lobbyName"]

	lobby, err := services.Lobby.GetLobby(lobbyName)
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

func LobbyStartHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lobbyName := vars["lobbyName"]

	lobby, err := services.Lobby.GetLobby(lobbyName)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	errors := services.CreateGame(lobby)
	if len(errors) > 0 {
		for _, err := range errors {
			log.Println(err)
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type submitCardsReqBody struct {
	CardIds []int `json:"cardIds"`
}

func CardsSubmitHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameName := vars["gameName"]
	playerName := vars["playerName"]
	tick, err := strconv.Atoi(vars["tick"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var reqBody submitCardsReqBody
	err = json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	errors := services.SubmitCards(gameName, playerName, tick, reqBody.CardIds)
	if len(errors) > 0 {
		for _, err := range errors {
			log.Println(err)
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type getTickResBody struct {
	Tick int `json:"tick"`
}

func GetCurrentTickHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameName := vars["gameName"]

	game, err := services.Game.GetGame(gameName)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resBody := getTickResBody{
		Tick: game.CurrentState.Tick,
	}
	err = json.NewEncoder(w).Encode(&resBody)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type gameStateResBody struct {
	Tick  int             `json:"tick"`
	Start state.GameState `json:"start"`
	End   state.GameState `json:"end"`
	// PlayersHand  []cards.Card    `json:"playersHand"`
}

func GetGameStateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameName := vars["gameName"]
	// playerName := vars["playerName"]
	tick, err := strconv.Atoi(vars["tick"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	game, err := services.Game.GetGame(gameName)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if game.CurrentState.Tick != tick {
		err = fmt.Errorf("tick %v does not match %v for game %s", tick, game.CurrentState.Tick, game.Name)
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO Implement player's hand only
	// player, err := game.GetPlayer(playerName)
	// if err != nil {
	// 	log.Println(err)
	// 	w.WriteHeader(http.StatusNotFound)
	// 	return
	// }

	resBody := gameStateResBody{
		// Note: Tick is for previous turn
		Tick:  game.PreviousState.Tick,
		Start: game.PreviousState,
		End:   game.CurrentState,
	}
	err = json.NewEncoder(w).Encode(&resBody)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
