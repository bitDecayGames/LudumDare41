package cards

import (
	"testing"
)

func TestUniquePriority(t *testing.T) {
	cardSet := CardSet{
		Cards: []Card{
			Card{ID: 1},
			Card{ID: 2},
			Card{ID: 3},
			Card{ID: 4},
		},
	}

	deckOne := NewDeckFromSet(cardSet, 2, 1)
	if len(deckOne) != 4 {
		t.Fatal("Deck does not contain 4 cards")
	}
	t.Logf("Done One of Two: %v", deckOne)

	deckTwo := NewDeckFromSet(cardSet, 2, 2)
	if len(deckTwo) != 4 {
		t.Fatal("Deck does not contain 4 cards")
	}
	t.Logf("Done Two of Two: %v", deckTwo)

	for _, cOne := range deckOne {
		for _, cTwo := range deckTwo {
			if cOne.Priority == cTwo.Priority {
				t.Fatalf("Duplicate priority between decks: %v", cOne.Priority)
			}
		}
	}
}
