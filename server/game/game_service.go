package game

import (
	"fmt"
	"sync"

	"github.com/bitDecayGames/LudumDare41/server/utils"

	"github.com/bitDecayGames/LudumDare41/server/cards"
	"github.com/bitDecayGames/LudumDare41/server/gameboard"
	"github.com/bitDecayGames/LudumDare41/server/lobby"
	"github.com/bitDecayGames/LudumDare41/server/state"
)

var mutex = &sync.Mutex{}

type GameService interface {
	NewGame(*lobby.Lobby, gameboard.GameBoard, cards.CardSet) *Game
	GetGame(string) (*Game, error)
}

type gameService struct {
	activeGames []*Game
}

func NewGameService() GameService {
	return &gameService{
		activeGames: []*Game{},
	}
}

func (gs *gameService) NewGame(lobby *lobby.Lobby, board gameboard.GameBoard, cardSet cards.CardSet) *Game {
	players := make(map[string]*state.Player)
	for _, player := range lobby.Players {
		players[player] = &state.Player{
			Name:    player,
			Hand:    make([]cards.Card, 0),
			Discard: make([]cards.Card, 0),
			Deck:    make([]cards.Card, 0),
			Pos: utils.Vector{
				X: 2,
				Y: 2,
			},
			Facing: utils.Vector{
				X: 0,
				Y: 1,
			},
		}
	}

	game := newGame(players, board, cardSet, lobby.Name)

	mutex.Lock()
	gs.activeGames = append(gs.activeGames, game)
	mutex.Unlock()

	return game
}

func (gs *gameService) GetGame(name string) (*Game, error) {
	for _, game := range gs.activeGames {
		if game.Name == name {
			return game, nil
		}
	}

	return nil, fmt.Errorf("game not found for name %s", name)
}
