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

	"github.com/bitDecayGames/LudumDare41/server/game"
	"github.com/bitDecayGames/LudumDare41/server/state"

	"github.com/bitDecayGames/LudumDare41/server/cards"
	"github.com/bitDecayGames/LudumDare41/server/gameboard"

	"github.com/bitDecayGames/LudumDare41/server/lobby"

	"github.com/bitDecayGames/LudumDare41/server/pubsub"
	"github.com/gorilla/mux"
)

// TODO Turn timeout logic

const (
	// Networking
	port        = 8080
	apiv1       = "/api/v1"
	testRoute   = apiv1 + "/test"
	pubsubRoute = apiv1 + "/pubsub"
	lobbyRoute  = apiv1 + "/lobby"
	gameRoute   = apiv1 + "/game"

	// Game
	minNumPlayers = 2
	maxNumPlayers = 4
)

type Services struct {
	pubsub pubsub.PubSubService
	lobby  lobby.LobbyService
	game   game.GameService
}

func NewServices() *Services {
	return &Services{
		pubsub: pubsub.NewPubSubService(),
		lobby:  lobby.NewLobbyService(),
		game:   game.NewGameService(),
	}
}

var services *Services

func main() {
	host := fmt.Sprintf(":%v", port)
	log.Printf("Starting server on %s ...", host)

	rand.Seed(time.Now().UnixNano())

	services = NewServices()

	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	// Tests
	r.HandleFunc(testRoute+"/ping", TestPingHandler).Methods("POST")
	r.HandleFunc(testRoute+"/game/{gameName}", TestCreateGameHandler).Methods("POST")
	r.HandleFunc(testRoute+"/game/{gameName}/player/{playerName}/cards", TestSubmitHandHandler).Methods("POST")
	// PubSub
	r.HandleFunc(pubsubRoute, PubSubHandler)
	r.HandleFunc(pubsubRoute+"/connection/{connectionID}", UpdatePubSubConnectionHandler).Methods("PUT")
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
		if reqURI == pubsubRoute {
			log.Printf("Skipping logging for %s", pubsubRoute)
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

type pingBody struct {
	GameName string `json:"gameName"`
}

type PingMessage struct {
	Status string `json:"status"`
}

func TestPingHandler(w http.ResponseWriter, r *http.Request) {
	var body pingBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	msg := pubsub.Message{
		MessageType: pubsub.PingMessage,
	}

	// TODO Change game name passed in?
	errors := services.pubsub.SendMessage("test", msg)
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

type testCreateGameReqBody struct {
	playerNames []string `json:"playeNames"`
}

func TestCreateGameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameName := vars["gameName"]

	var reqBody testCreateGameReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	lobby, err := services.lobby.NewLobby()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO Might need a mutex lock here
	lobby.Name = gameName

	for _, playerName := range reqBody.playerNames {
		_, err = lobby.AddPlayer(playerName)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	errors := CreateGame(lobby)
	if len(errors) > 0 {
		for _, err := range errors {
			log.Println(err)
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func TestSubmitHandHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameName := vars["gameName"]
	playerName := vars["playerName"]

	game, err := services.game.GetGame(gameName)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	player, err := game.GetPlayer(playerName)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// For now, always grab the first 3 cards
	cardIds := []int{}
	for _, card := range player.Hand[0:3] {
		cardIds = append(cardIds, card.ID)
	}

	errors := SubmitCards(gameName, playerName, game.CurrentState.Tick, cardIds)
	if len(errors) > 0 {
		for _, err := range errors {
			log.Println(err)
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func PubSubHandler(w http.ResponseWriter, r *http.Request) {
	connectionID, err := services.pubsub.AddSubscription(w, r)
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

	err = services.pubsub.UpdateSubscription(connectionID, body.GameName, body.PlayerName)
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
	lobby, err := services.lobby.NewLobby()
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

	lobby, err := services.lobby.GetLobby(lobbyName)
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
	errors := services.pubsub.SendMessage(lobbyName, msg)
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

	lobby, err := services.lobby.GetLobby(lobbyName)
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

	lobby, err := services.lobby.GetLobby(lobbyName)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	errors := CreateGame(lobby)
	if len(errors) > 0 {
		for _, err := range errors {
			log.Println(err)
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func CreateGame(lobby *lobby.Lobby) []error {
	// TODO Allow different boards and card sets.
	board := gameboard.LoadBoard("default")
	cardSet := cards.LoadSet("default")
	game := services.game.NewGame(lobby, board, cardSet)

	// TODO Fix
	// if len(game.Players) < minNumPlayers {
	// 	err := fmt.Errorf("minimum number of %v players not met: %v", minNumPlayers, game.Players)
	// 	return []error{err}
	// }

	// if len(game.Players) > maxNumPlayers {
	// 	err := fmt.Errorf("maximum number of %v players exceeded: %v", maxNumPlayers, game.Players)
	// 	return []error{err}
	// }

	msg := pubsub.Message{
		MessageType: pubsub.GameUpdateMessage,
		ID:          game.Name,
		Tick:        game.CurrentState.Tick,
	}
	return services.pubsub.SendMessage(game.Name, msg)
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

	errors := SubmitCards(gameName, playerName, tick, reqBody.CardIds)
	if len(errors) > 0 {
		for _, err := range errors {
			log.Println(err)
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func SubmitCards(gameName, playerName string, tick int, cardIds []int) []error {
	game, err := services.game.GetGame(gameName)
	if err != nil {
		return []error{err}
	}

	err = game.SubmitCards(playerName, tick, cardIds)
	if err != nil {
		return []error{err}
	}

	// Check for advance to next turn
	if game.AreSubmissionsComplete() {
		log.Printf("Starting next turn for game %s at tick %v", game.Name, game.CurrentState.Tick)

		_ = game.AggregateTurn()
		game.ExecuteTurn()

		log.Printf("Turn complete for game %s at tick %v", game.Name, game.CurrentState.Tick)

		msg := pubsub.Message{
			MessageType: pubsub.GameUpdateMessage,
			ID:          game.Name,
			Tick:        game.CurrentState.Tick,
		}
		return services.pubsub.SendMessage(game.Name, msg)
	}

	return []error{}
}

type getTickResBody struct {
	Tick int `json:"tick"`
}

func GetCurrentTickHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameName := vars["gameName"]

	game, err := services.game.GetGame(gameName)
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
	PlayersHand  []cards.Card    `json:"playersHand"`
	CurrentState state.GameState `json:"currentState"`
}

func GetGameStateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameName := vars["gameName"]
	playerName := vars["playerName"]
	tick, err := strconv.Atoi(vars["tick"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	game, err := services.game.GetGame(gameName)
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

	player, err := game.GetPlayer(playerName)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resBody := gameStateResBody{
		PlayersHand:  player.Hand,
		CurrentState: game.CurrentState,
	}
	err = json.NewEncoder(w).Encode(&resBody)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
