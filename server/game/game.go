package game

import (
	"github.com/bitDecayGames/LudumDare41/server/cards"
	"github.com/bitDecayGames/LudumDare41/server/gameboard"
	"github.com/bitDecayGames/LudumDare41/server/lobby"
)

type GameService interface {
	NewGame(*lobby.Lobby, gameboard.GameBoard, cards.CardSet) Game
}

type Player string

type Game struct {
	Players []Player
	Board   gameboard.GameBoard
}

type gameService struct {
	activeGames []Game
}

func NewGameService() GameService {
	return &gameService{}
}

func (gs *gameService) NewGame(lobby *lobby.Lobby, board gameboard.GameBoard, cardSet cards.CardSet) Game {
	return Game{}
}
