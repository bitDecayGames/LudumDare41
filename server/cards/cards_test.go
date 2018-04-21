package cards

import "testing"

func TestCardSet(t *testing.T) {
	cardSet := LoadSet("default")
	if len(cardSet.Cards) <= 0 {
		t.Error("No cards in the deck")
	}
}
