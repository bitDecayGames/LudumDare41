package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/bitDecayGames/LudumDare41/server/logic"

	"github.com/bitDecayGames/LudumDare41/server/state"
	"github.com/gorilla/mux"
)

const (
	gameRoute = apiv1 + "/game"
)

type GameRoutes struct {
	services *Services
}

func (gr *GameRoutes) AddRoutes(r *mux.Router) {
	r.HandleFunc(gameRoute+"/{gameName}/tick/{tick}/player/{playerName}/cards", gr.cardsSubmitHandler).Methods("PUT")
	r.HandleFunc(gameRoute+"/{gameName}/tick", gr.getCurrentTickHandler).Methods("GET")
	r.HandleFunc(gameRoute+"/{gameName}/tick/{tick}/player/{playerName}", gr.getGameStateHandler).Methods("GET")
}

type submitCardsReqBody struct {
	CardIds []int `json:"cardIds"`
}

func (gr *GameRoutes) cardsSubmitHandler(w http.ResponseWriter, r *http.Request) {
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

	errors := gr.services.SubmitCards(gameName, playerName, tick, reqBody.CardIds)
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

func (gr *GameRoutes) getCurrentTickHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameName := vars["gameName"]

	game, err := gr.services.Game.GetGame(gameName)
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
	Tick  int                `json:"tick"`
	Start state.GameState    `json:"start"`
	End   state.GameState    `json:"end"`
	Diff  logic.StepSequence `json:"diff"`
	// PlayersHand  []cards.Card    `json:"playersHand"`
}

func (gr *GameRoutes) getGameStateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameName := vars["gameName"]
	// playerName := vars["playerName"]
	tick, err := strconv.Atoi(vars["tick"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	game, err := gr.services.Game.GetGame(gameName)
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
		Diff:  game.LastSequence,
	}
	err = json.NewEncoder(w).Encode(&resBody)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func NewGameRoutes(services *Services) *GameRoutes {
	return &GameRoutes{
		services: services,
	}
}
