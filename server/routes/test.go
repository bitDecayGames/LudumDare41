package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bitDecayGames/LudumDare41/server/pubsub"
	"github.com/gorilla/mux"
)

const (
	testRoute = apiv1 + "/test"
)

type TestRoutes struct {
	services *Services
}

func (tr *TestRoutes) AddRoutes(r *mux.Router) {
	r.HandleFunc(testRoute+"/ping", tr.testPingHandler).Methods("POST")
	r.HandleFunc(testRoute+"/game/{gameName}", tr.testCreateGameHandler).Methods("POST")
	r.HandleFunc(testRoute+"/game/{gameName}/player/{playerName}/cards", tr.testSubmitHandHandler).Methods("POST")
}

type pingBody struct {
	GameName string `json:"gameName"`
}

type PingMessage struct {
	Status string `json:"status"`
}

func (tr *TestRoutes) testPingHandler(w http.ResponseWriter, r *http.Request) {
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
	errors := tr.services.PubSub.SendMessage("test", msg)
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

func (tr *TestRoutes) testCreateGameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameName := vars["gameName"]

	var reqBody testCreateGameReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	lobby, err := tr.services.Lobby.NewLobby()
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

	errors := tr.services.CreateGame(lobby)
	if len(errors) > 0 {
		for _, err := range errors {
			log.Println(err)
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (tr *TestRoutes) testSubmitHandHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameName := vars["gameName"]
	playerName := vars["playerName"]

	game, err := tr.services.Game.GetGame(gameName)
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

	errors := tr.services.SubmitCards(gameName, playerName, game.CurrentState.Tick, cardIds)
	if len(errors) > 0 {
		for _, err := range errors {
			log.Println(err)
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func NewTestRoutes(services *Services) *TestRoutes {
	return &TestRoutes{
		services: services,
	}
}
