package logic

import (
	"testing"

	"github.com/bitDecayGames/LudumDare41/server/cards"
	"github.com/bitDecayGames/LudumDare41/server/state"
	"github.com/bitDecayGames/LudumDare41/server/utils"
)

func TestApplyingCard(t *testing.T) {
	testCard := CardTypeMap[Card_move_forward_1]
	testCard.Owner = "player1"
	testPlayer := state.Player{
		Name:    "player1",
		Discard: make([]cards.Card, 0),
		Hand:    []cards.Card{testCard},
		Pos:     utils.Vector{0, 0},
		Facing:  utils.Vector{0, 1},
	}
	gs := state.GameState{
		Players: []state.Player{testPlayer},
	}

	_, newState := ApplyCard(testCard, gs)
	if len(newState.Players[0].Hand) != 0 {
		t.Fatal("Card was not discarded after use")
	}

	if newState.Players[0].Pos.X != 0 && newState.Players[0].Pos.Y != 1 {
		t.Fatal("Player was not properly moved forward")
	}
}
