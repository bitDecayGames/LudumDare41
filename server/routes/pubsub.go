package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	PubSubRoute = apiv1 + "/pubsub"
)

type PubSubRoutes struct {
	services *Services
}

func (psr *PubSubRoutes) AddRoutes(r *mux.Router) {
	r.HandleFunc(PubSubRoute, psr.pubSubHandler)
	r.HandleFunc(PubSubRoute+"/connection/{connectionID}", psr.updatePubSubConnectionHandler).Methods("PUT")
}

func (psr *PubSubRoutes) pubSubHandler(w http.ResponseWriter, r *http.Request) {
	connectionID, err := psr.services.PubSub.AddSubscription(w, r)
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

func (psr *PubSubRoutes) updatePubSubConnectionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	connectionID := vars["connectionID"]

	var body updateSubBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = psr.services.PubSub.UpdateSubscription(connectionID, body.GameName, body.PlayerName)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func NewPubSubRoutes(services *Services) *PubSubRoutes {
	return &PubSubRoutes{
		services: services,
	}
}
