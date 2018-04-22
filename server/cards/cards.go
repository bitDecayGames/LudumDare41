package cards

import (
	"math/rand"
)

// Card represents a single unique playing card
type Card struct {
	ID       int    `json:"id"`
	Priority int    `json:"priority"`
	Owner    string `json:"owner"`
	CardType string `json:"cardType"`
}

// CardSet is an instance of a set of cards
type CardSet struct {
	Cards []Card
}

func LoadSet(setDef map[string]int) CardSet {
	set := CardSet{
		Cards: []Card{},
	}

	for cType, num := range setDef {
		for ;num > 0;num-- {
			set.Cards = append(set.Cards, Card{
				CardType:cType,
			})
		}
	}

	return set
}

// NewDeckFromSet will generate a new deck of cards with a unique priority assuming only a `playerCount` number of decks are generated and with a unique set of `playerNumber`
// `playerNumber` should be one-based
func NewDeckFromSet(cardSet CardSet, playerCount int, playerNumber int) []Card {
	deck := make([]Card, len(cardSet.Cards))
	perm := rand.Perm(len(cardSet.Cards))
	for i, randIndex := range perm {
		deck[i] = cardSet.Cards[randIndex]
		// This should ensure unique priorities across a `playerCount` number of decks
		priority := (randIndex * playerCount) + (playerNumber - 1)
		deck[i].Priority = priority
		deck[i].ID = priority
	}

	return deck
}

func ShuffleCards(inCards []Card) []Card {
	shuffled := make([]Card, len(inCards))
	perm := rand.Perm(len(inCards))
	for i, randIndex := range perm {
		shuffled[i] = inCards[randIndex]
	}

	return shuffled
}
