package game

import (
	"github.com/bitDecayGames/LudumDare41/server/gameboard"
)

type Player string

type Game struct {
	Players []Player
	Board   gameboard.GameBoard
}

type GameService struct {
	activeGames []Game
}

func NewGameService() *GameService {

}

func (gs *GameService) NewGame() {

}
