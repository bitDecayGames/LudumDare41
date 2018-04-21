package game

import (
	"fmt"
	"sort"

	"github.com/bitDecayGames/LudumDare41/server/cards"
	"github.com/bitDecayGames/LudumDare41/server/gameboard"
)

const HAND_SIZE = 5

type Game struct {
	Players map[string]*Player
	Board   gameboard.GameBoard

	pendingSubmissions map[string][]cards.Card
}

func newGame(players map[string]*Player, board gameboard.GameBoard) *Game {
	return &Game{
		Players:            players,
		Board:              board,
		pendingSubmissions: make(map[string][]cards.Card),
	}
}

func (g *Game) DealCards() {
	priority := 1
	for _, player := range g.Players {
		for len(player.Hand) < HAND_SIZE {
			// TODO: Actually pull cards from the player deck. Shuffle cards if needed
			player.Hand = append(player.Hand, cards.Card{ID: 1, Priority: priority})
			priority += 1
		}
	}
}

func (g *Game) SubmitCards(player string, cards []cards.Card) error {
	// TODO validate these cards
	if g.pendingSubmissions[player] != nil {
		return fmt.Errorf("Player already has a pending submission")
	}

	g.pendingSubmissions[player] = cards
	return nil
}

func (g *Game) AggregateTurn() []cards.Card {
	cardOrder := make([]cards.Card, 0)
	for _, pendingCards := range g.pendingSubmissions {
		cardOrder = append(cardOrder, pendingCards...)
	}
	g.pendingSubmissions = make(map[string][]cards.Card)

	sort.Slice(cardOrder, func(i, j int) bool { return cardOrder[i].Priority > cardOrder[j].Priority })
	return cardOrder
}
