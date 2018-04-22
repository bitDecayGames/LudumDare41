package gameboard

import (
	"testing"
)

func TestMapLoading(t *testing.T) {
	board := LoadBoard("default")
	if board.Tiles == nil {
		t.Error("game board has nil tiles")
	}

	if len(board.Tiles) <= 0 {
		t.Error("Game board has no width")
	}

	if len(board.Tiles[0]) <= 0 {
		t.Error("Game board has no height")
	}

	flattendTileCount := 0

	for x, yTiles := range board.Tiles {
		for y, tile := range yTiles {
			if tile.Pos.X != x {
				t.Errorf("tile position x: %v does not match x: %v", tile.Pos.X, x)
			}

			if tile.Pos.Y != y {
				t.Errorf("tile position y: %v does not match y: %v", tile.Pos.Y, y)
			}

			flattendTileCount++
		}
	}

	if len(board.FlattenedTiles) != flattendTileCount {
		t.Errorf("flattenedTiles length was %v, expected %v", len(board.FlattenedTiles), flattendTileCount)
	}
}
