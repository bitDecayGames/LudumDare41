package lobby

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/satori/go.uuid"
	"github.com/speps/go-hashids"
)

const (
	defaultAlphabet    = "acdefghjkmnpqrstuvwxyz2345789"
	lobbyNameMinLength = 6
	randRange          = 999999
)

var mutex = &sync.Mutex{}

type LobbyService interface {
	NewLobby() (*Lobby, error)
	GetLobbies() []*Lobby
}

type lobbyService struct {
	lobbies []*Lobby
	hashID  *hashids.HashID
}

type Lobby struct {
	Name    string
	Players []string
}

func (ls *lobbyService) NewLobby() (*Lobby, error) {
	lobbyName, err := ls.genLobbyName()
	if err != nil {
		return nil, err
	}

	lobby := &Lobby{
		Name:    lobbyName,
		Players: []string{},
	}

	mutex.Lock()
	ls.lobbies = append(ls.lobbies, lobby)
	mutex.Unlock()

	return lobby, nil
}

func (ls *lobbyService) GetLobbies() []*Lobby {
	return ls.lobbies
}

func NewLobbyService() LobbyService {
	hd := hashids.NewData()
	hd.Alphabet = defaultAlphabet
	hd.MinLength = lobbyNameMinLength
	hd.Salt = uuid.NewV4().String()
	hashID := hashids.NewWithData(hd)

	return &lobbyService{
		lobbies: []*Lobby{},
		hashID:  hashID,
	}
}

func (l *Lobby) AddPlayer(name string) error {
	// Check for exisiting player
	for _, n := range l.Players {
		if name == n {
			return fmt.Errorf("Player with name %s already exists", name)
		}
	}

	mutex.Lock()
	l.Players = append(l.Players, name)
	mutex.Unlock()

	return nil
}

func (l *Lobby) GetPlayers() []string {
	return l.Players
}

func (ls *lobbyService) genLobbyName() (string, error) {
	lobbyName, err := ls.hashID.Encode([]int{rand.Intn(randRange)})
	return lobbyName, err
}
