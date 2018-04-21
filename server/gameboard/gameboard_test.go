package gameboard

import "testing"

func TestMapLoading(t *testing.T) {
	board := LoadBoard("default")
	if board.tiles == nil {
		t.Error("game board has nil tiles")
	}

	if len(board.tiles) <= 0 {
		t.Error("Game board has no width")
	}

	if len(board.tiles[0]) <= 0 {
		t.Error("Game board has no height")
	}
}
