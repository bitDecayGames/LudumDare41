package state

import (
	"testing"

	"github.com/bitDecayGames/LudumDare41/server/cards"
)

func TestPlayerHandFuncs(t *testing.T) {
	c1 := cards.Card{ID: 1, Owner: "test"}
	c2 := cards.Card{ID: 2, Owner: "test"}
	c3 := cards.Card{ID: 3, Owner: "test"}
	c4 := cards.Card{ID: 4, Owner: "test"}

	p := &Player{
		Name: "test",
		Hand: []cards.Card{c1, c2, c3, c4},
	}

	p.DiscardCard(c3)

	if len(p.Hand) != 3 || len(p.Discard) != 1 {
		t.Fatal("Card was not discarded properly")
	}

	p.DiscardEntireHand()

	if len(p.Hand) != 0 || len(p.Discard) != 4 {
		t.Fatal("Entire hand was not discarded properly")
	}
}
