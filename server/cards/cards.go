package cards

import (
	"math/rand"
)

// Card represents a single unique playing card
type Card struct {
	ID       int
	Priority int
	Owner    string
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

// NewDeckFromSet will generate a new deck of cards with a unique priority assuming only a `playerCount` number of decks are generated and with a unique set of `playerNumber`
// `playerNumber` should be one-based
func NewDeckFromSet(cardSet CardSet, playerCount int, playerNumber int) []Card {
	deck := make([]Card, len(cardSet.Cards))
	perm := rand.Perm(len(cardSet.Cards))
	for i, randIndex := range perm {
		deck[i] = cardSet.Cards[randIndex]
		// This should ensure unique priorities across a `playerCount` number of decks
		deck[i].Priority = (randIndex * playerCount) + (playerNumber - 1)
	}

	return deck
}
