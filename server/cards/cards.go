package cards

// Card represents a single unique playing card
type Card struct {
	ID int
}

// CardSet is an instance of a set of cards
type CardSet struct {
	Cards []Card
}

func LoadSet(name string) CardSet {
	return CardSet{
		Cards: []Card{
			Card{ID: 1},
		},
	}
}
