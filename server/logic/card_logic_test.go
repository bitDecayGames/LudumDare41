package logic

import (
	"testing"

	"github.com/bitDecayGames/LudumDare41/server/gameboard"

	"github.com/bitDecayGames/LudumDare41/server/cards"
	"github.com/bitDecayGames/LudumDare41/server/state"
	"github.com/bitDecayGames/LudumDare41/server/utils"
)

func TestApplyingCardMoveForward(t *testing.T) {
	testCard := CardTypeMap[Card_move_forward_1]
	testCard.Owner = "player1"
	testPlayer := state.Player{
		Name:    "player1",
		Discard: make([]cards.Card, 0),
		Hand:    []cards.Card{testCard},
		Pos:     utils.Vector{X: 0, Y: 0},
		Facing:  utils.Vector{X: 0, Y: 1},
	}
	gs := state.GameState{
		Players: []state.Player{testPlayer},
		Board: gameboard.GameBoard{
			Tiles: [][]gameboard.Tile{[]gameboard.Tile{gameboard.Tile{TileType: gameboard.Empty_tile}, gameboard.Tile{TileType: gameboard.Empty_tile}}},
		},
	}

	seq, newState := ApplyCard(testCard, gs)
	if len(newState.Players[0].Hand) != 0 {
		t.Fatal("Card was not discarded after use")
	}

	if len(seq.steps) != 1 {
		t.Fatalf("Sequence not proper: %v", seq)
	}

	if !(newState.Players[0].Pos.X == 0 && newState.Players[0].Pos.Y == 1) {
		t.Fatal("Player was not properly moved forward")
	}
}

func TestApplyingCardMoveForwardPartialBlock(t *testing.T) {
	testCard := CardTypeMap[Card_move_forward_3]
	testCard.Owner = "player1"
	testPlayer := state.Player{
		Name:    "player1",
		Discard: make([]cards.Card, 0),
		Hand:    []cards.Card{testCard},
		Pos:     utils.Vector{X: 0, Y: 0},
		Facing:  utils.Vector{X: 0, Y: 1},
	}
	gs := state.GameState{
		Players: []state.Player{testPlayer},
		Board: gameboard.GameBoard{
			Tiles: [][]gameboard.Tile{
				[]gameboard.Tile{
					gameboard.Tile{TileType: gameboard.Empty_tile},
					gameboard.Tile{TileType: gameboard.Empty_tile},
					gameboard.Tile{TileType: gameboard.Empty_tile},
				},
			},
		},
	}

	seq, newState := ApplyCard(testCard, gs)
	if len(newState.Players[0].Hand) != 0 {
		t.Fatal("Card was not discarded after use")
	}

	if len(seq.steps) != 2 {
		t.Fatalf("Sequence not proper: %v", seq)
	}

	if !(newState.Players[0].Pos.X == 0 && newState.Players[0].Pos.Y == 2) {
		t.Fatal("Player was not properly moved forward")
	}
}

func TestApplyingCardMoveBackward(t *testing.T) {
	testCard := CardTypeMap[Card_move_backward_1]
	testCard.Owner = "player1"
	testPlayer := state.Player{
		Name:    "player1",
		Discard: make([]cards.Card, 0),
		Hand:    []cards.Card{testCard},
		Pos:     utils.Vector{X: 0, Y: 0},
		Facing:  utils.Vector{X: 0, Y: -1},
	}
	gs := state.GameState{
		Players: []state.Player{testPlayer},
		Board: gameboard.GameBoard{
			Tiles: [][]gameboard.Tile{[]gameboard.Tile{gameboard.Tile{TileType: gameboard.Empty_tile}, gameboard.Tile{TileType: gameboard.Empty_tile}}},
		},
	}

	seq, newState := ApplyCard(testCard, gs)
	if len(newState.Players[0].Hand) != 0 {
		t.Fatal("Card was not discarded after use")
	}

	if len(seq.steps) != 1 {
		t.Fatalf("Sequence not proper: %v", seq)
	}

	if !(newState.Players[0].Pos.X == 0 && newState.Players[0].Pos.Y == 1) {
		t.Fatal("Player was not properly moved forward")
	}
}

func TestPushingPlayer(t *testing.T) {
	testCard := CardTypeMap[Card_move_forward_1]
	testCard.Owner = "player1"
	testPlayer := state.Player{
		Name:    "player1",
		Discard: make([]cards.Card, 0),
		Hand:    []cards.Card{testCard},
		Pos:     utils.Vector{X: 0, Y: 0},
		Facing:  utils.Vector{X: 0, Y: 1},
	}
	testPlayerTwo := state.Player{
		Name:    "player2",
		Discard: make([]cards.Card, 0),
		Hand:    []cards.Card{testCard},
		Pos:     utils.Vector{X: 0, Y: 1},
		Facing:  utils.Vector{X: 1, Y: 0},
	}

	gs := state.GameState{
		Players: []state.Player{testPlayer, testPlayerTwo},
		Board: gameboard.GameBoard{
			Tiles: [][]gameboard.Tile{
				[]gameboard.Tile{
					gameboard.Tile{TileType: gameboard.Empty_tile},
					gameboard.Tile{TileType: gameboard.Empty_tile},
					gameboard.Tile{TileType: gameboard.Empty_tile},
				},
			},
		},
	}

	seq, newState := ApplyCard(testCard, gs)
	if len(newState.Players[0].Hand) != 0 {
		t.Fatal("Card was not discarded after use")
	}

	if len(seq.steps) != 1 {
		t.Fatalf("Sequence not proper: %v", seq)
	}

	if len(seq.steps[0].actions) != 2 {
		t.Fatal("Step did not have two actions")
	}

	if newState.Players[0].Pos.X != 0 && newState.Players[0].Pos.Y != 1 {
		t.Fatal("Player one was not properly moved forward")
	}

	if newState.Players[1].Pos.X != 0 && newState.Players[1].Pos.Y != 2 {
		t.Fatal("Player  two was not properly moved forward")
	}
}

func TestApplyingCardRotate(t *testing.T) {
	testCard := CardTypeMap[Card_rotate_clockwise]
	testCard.Owner = "player1"
	testPlayer := state.Player{
		Name:    "player1",
		Discard: make([]cards.Card, 0),
		Hand:    []cards.Card{testCard},
		Pos:     utils.Vector{X: 0, Y: 0},
		Facing:  utils.Vector{X: 0, Y: 1},
	}
	gs := state.GameState{
		Players: []state.Player{testPlayer},
		// Board: gameboard.GameBoard{
		// 	Tiles: [][]gameboard.Tile{[]gameboard.Tile{gameboard.Tile{TileType: gameboard.Empty_tile}, gameboard.Tile{TileType: gameboard.Empty_tile}}},
		// },
	}

	seq, newState := ApplyCard(testCard, gs)
	if len(newState.Players[0].Hand) != 0 {
		t.Fatal("Card was not discarded after use")
	}

	if len(seq.steps) != 1 {
		t.Fatalf("Sequence not proper: %v", seq)
	}

	expected := utils.Vector{X: 1, Y: 0}
	if !(utils.VecEquals(newState.Players[0].Facing, expected)) {
		t.Fatalf("Player was not properly rotated. Expected %v, got %v", expected, newState.Players[0].Facing)
	}
}

func TestShootingPlayer(t *testing.T) {
	testCard := CardTypeMap[Card_shoot_main_turret]
	testCard.Owner = "player1"
	testPlayer := state.Player{
		Name:    "player1",
		Discard: make([]cards.Card, 0),
		Hand:    []cards.Card{testCard},
		Pos:     utils.Vector{X: 0, Y: 0},
		Facing:  utils.Vector{X: 0, Y: 1},
	}
	testPlayerTwo := state.Player{
		Name:    "player2",
		Discard: make([]cards.Card, 0),
		Hand:    []cards.Card{testCard},
		Pos:     utils.Vector{X: 0, Y: 4},
		Facing:  utils.Vector{X: 1, Y: 0},
	}

	gs := state.GameState{
		Players: []state.Player{testPlayer, testPlayerTwo},
		Board: gameboard.GameBoard{
			Tiles: [][]gameboard.Tile{
				[]gameboard.Tile{
					gameboard.Tile{TileType: gameboard.Empty_tile},
					gameboard.Tile{TileType: gameboard.Empty_tile},
					gameboard.Tile{TileType: gameboard.Empty_tile},
					gameboard.Tile{TileType: gameboard.Empty_tile},
					gameboard.Tile{TileType: gameboard.Empty_tile},
					gameboard.Tile{TileType: gameboard.Empty_tile},
				},
			},
		},
	}

	seq, newState := ApplyCard(testCard, gs)
	if len(newState.Players[0].Hand) != 0 {
		t.Fatal("Card was not discarded after use")
	}

	if len(seq.steps) != 2 {
		// expecting shoot, die
		t.Fatalf("Sequence not proper: %v", seq)
	}

	if newState.Players[1].Pos.X != -1 && newState.Players[1].Pos.Y != -1 {
		t.Fatal("Player  two was not properly killed")
	}
}
