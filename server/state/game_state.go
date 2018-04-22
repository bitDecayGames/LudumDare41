package state

import (
	"github.com/bitDecayGames/LudumDare41/server/gameboard"
	"github.com/bitDecayGames/LudumDare41/server/utils"
)

type GameState struct {
	Tick    int                 `json:"tick"`
	Players []Player            `json:"players"`
	Crate   utils.Vector        `json:"crate"`
	Board   gameboard.GameBoard `json:"gameBoard"`
}

func NewState(tick int, players map[string]*Player, board gameboard.GameBoard) GameState {
	playersData := make([]Player, 0)
	for _, p := range players {
		playersData = append(playersData, *p)
	}

	return GameState{
		Tick:    tick,
		Players: playersData,
		Board:   board,
	}
}
