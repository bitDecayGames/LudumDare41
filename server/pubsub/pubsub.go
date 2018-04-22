package pubsub

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/satori/go.uuid"

	"github.com/gorilla/websocket"
)

const (
	PingMessage              = "ping"
	ConnectionStartedMessage = "connectionStarted"
	PlayerJoinMessage        = "playerJoin"
	GameUpdateMessage        = "gameUpdate"
)

type Message struct {
	MessageType string `json:"messageType"`
	ID          string `json:"id"`
	// TODO Find a better way of adding custom fields without bloating this struct.
	Tick int `json:"tick"`
}

var upgrader = websocket.Upgrader{
	// TODO Remove to prevent cross origin requests
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var mutex = &sync.Mutex{}

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

func (ps *pubSubService) AddSubscription(w http.ResponseWriter, r *http.Request) (string, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return "", err
	}

	sub := &subscription{
		connectionID: uuid.NewV4().String(),
		conn:         conn,
	}

	mutex.Lock()
	ps.subscriptions = append(ps.subscriptions, sub)
	mutex.Unlock()

	body := Message{
		MessageType: ConnectionStartedMessage,
		ID:          sub.connectionID,
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

	mutex.Lock()
	curSub.gameName = gameName
	curSub.playerName = playerName
	mutex.Unlock()

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
			log.Printf("Sent message %+v to game %s", gameName)
		}
	}

	return errors
}

func NewPubSubService() PubSubService {
	return &pubSubService{
		subscriptions: []*subscription{},
	}
}
