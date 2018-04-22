package game

import (
	"github.com/bitDecayGames/LudumDare41/server/cards"
	"github.com/bitDecayGames/LudumDare41/server/gameboard"
	"github.com/bitDecayGames/LudumDare41/server/lobby"
	"github.com/bitDecayGames/LudumDare41/server/state"
)

type GameService interface {
	NewGame(*lobby.Lobby, gameboard.GameBoard, cards.CardSet) *Game
}

type gameService struct {
	activeGames []Game
}

func NewGameService() GameService {
	return &gameService{}
}

func (gs *gameService) NewGame(lobby *lobby.Lobby, board gameboard.GameBoard, cardSet cards.CardSet) *Game {
	players := make(map[string]*state.Player)
	for _, player := range lobby.Players {
		players[player] = &state.Player{
			Name:    player,
			Hand:    make([]cards.Card, 0),
			Discard: make([]cards.Card, 0),
			Deck:    make([]cards.Card, 0),
		}
	}
	return newGame(players, board, cardSet)
}
