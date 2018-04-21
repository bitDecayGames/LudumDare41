package lobby

import "fmt"

type LobbyService interface {
	NewLobby() *Lobby
	GetLobbies() []*Lobby
}

type lobbyService struct {
	lobbies []*Lobby
}

type Lobby struct {
	Name    string
	Players []string
}

func (ls *lobbyService) NewLobby() *Lobby {
	// TODO Generate random, non-conflicting name
	lobbyName := "test"
	lobby := &Lobby{
		Name:    lobbyName,
		Players: []string{},
	}
	ls.lobbies = append(ls.lobbies, lobby)
	return lobby
}

func (ls *lobbyService) GetLobbies() []*Lobby {
	return ls.lobbies
}

func NewLobbyService() LobbyService {
	return &lobbyService{
		lobbies: []*Lobby{},
	}
}

func (l *Lobby) AddPlayer(name string) error {
	// Check for exisiting player
	for _, n := range l.Players {
		if name == n {
			return fmt.Errorf("Player with name %s already exists", name)
		}
	}

	l.Players = append(l.Players, name)
	fmt.Println(l.Players)
	return nil
}
