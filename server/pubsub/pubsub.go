package pubsub

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
	Type string `json:"type"`
}

var upgrader = websocket.Upgrader{}

type PubSubService interface {
	AddSubscription(http.ResponseWriter, *http.Request) error
	SendMessage(Message) error
}

type pubSubService struct {
	connections []*websocket.Conn
}

func (ps *pubSubService) AddSubscription(w http.ResponseWriter, r *http.Request) error {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	ps.connections = append(ps.connections, conn)
	return nil
}

func (ps *pubSubService) SendMessage(m Message) error {
	messageBytes, err := json.Marshal(m)
	if err != nil {
		return err
	}

	for _, conn := range ps.connections {
		err := conn.WriteMessage(websocket.TextMessage, messageBytes)
		if err != nil {
			// TODO remove connection?
			log.Println(err)
		}
	}

	return nil
}

func NewPubSubService() PubSubService {
	return &pubSubService{
		connections: []*websocket.Conn{},
	}
}
