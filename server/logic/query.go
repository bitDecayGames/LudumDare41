package logic

import (
	"github.com/bitDecayGames/LudumDare41/server/gameboard"
	"github.com/bitDecayGames/LudumDare41/server/state"
	"github.com/bitDecayGames/LudumDare41/server/utils"
)

func isEmptyTile(target utils.Vector, g state.GameState) bool {
	if target.X < 0 || target.X >= len(g.Board.Tiles) {
		return false
	}

	if target.Y < 0 || target.Y >= len(g.Board.Tiles[0]) {
		return false
	}

	return g.Board.Tiles[target.X][target.Y].TileType == gameboard.Empty_tile
}

func isPlayerOccupying(target utils.Vector, g state.GameState) bool {
	for _, p := range g.Players {
		if utils.VecEquals(p.Pos, target) {
			return true
		}
	}
	return false
}

func getPlayerAtPos(target utils.Vector, g state.GameState) *state.Player {
	for i, p := range g.Players {
		if utils.VecEquals(p.Pos, target) {
			return &g.Players[i]
		}
	}
	return &state.Player{}
}
