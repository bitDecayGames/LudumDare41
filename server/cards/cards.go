package cards

// Card represents a single unique playing card
type Card struct {
	ID int
}

// Deck is an instance of a set of cards
type Deck struct {
	Cards []Card
}

func LoadSet(name string) Deck {
	return Deck{
		Cards: []Card{
			Card{ID: 1},
		},
	}
}
