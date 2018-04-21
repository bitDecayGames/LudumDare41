package pubsub

import (
	"fmt"
	"log"
	"net/http"

	"github.com/satori/go.uuid"

	"github.com/gorilla/websocket"
)

type Message struct {
	Type string `json:"type"`
}

var upgrader = websocket.Upgrader{
	// TODO Remove to prevent cross origin requests
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type PubSubService interface {
	AddSubscription(http.ResponseWriter, *http.Request) (string, error)
	UpdateSubscription(connectionID string, gameName string, playerName string) error
	SendMessage(gameName string, msg Message) []error
}

type pubSubService struct {
	subscriptions []*subscription
}

type subscription struct {
	conn         *websocket.Conn
	connectionID string
	gameName     string
	playerName   string
}

type connectionBody struct {
	ConnectionID string `json:"connectionID"`
}

func (ps *pubSubService) AddSubscription(w http.ResponseWriter, r *http.Request) (string, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return "", err
	}

	sub := &subscription{
		connectionID: uuid.NewV4().String(),
		conn:         conn,
	}

	ps.subscriptions = append(ps.subscriptions, sub)

	body := connectionBody{
		ConnectionID: sub.connectionID,
	}

	return sub.connectionID, conn.WriteJSON(body)
}

func (ps *pubSubService) UpdateSubscription(connectionID, gameName, playerName string) error {
	var curSub *subscription
	for _, sub := range ps.subscriptions {
		if sub.connectionID == connectionID {
			curSub = sub
			break
		}
	}

	if curSub == nil {
		return fmt.Errorf("subscription not found for connectionID %s", connectionID)
	}

	curSub.gameName = gameName
	curSub.playerName = playerName

	return nil
}

func (ps *pubSubService) SendMessage(gameName string, msg Message) []error {
	var errors []error
	for _, sub := range ps.subscriptions {
		if sub.gameName == gameName {
			err := sub.conn.WriteJSON(msg)
			if err != nil {
				log.Printf("%s, %v, %v", gameName, msg, err)
				errors = append(errors, err)
			}
		}
	}

	return errors
}

func NewPubSubService() PubSubService {
	return &pubSubService{
		subscriptions: []*subscription{},
	}
}
