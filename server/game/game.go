package game

import (
	"fmt"
	"sort"

	"github.com/bitDecayGames/LudumDare41/server/cards"
	"github.com/bitDecayGames/LudumDare41/server/gameboard"
	"github.com/bitDecayGames/LudumDare41/server/logic"
	"github.com/bitDecayGames/LudumDare41/server/state"
)

const HAND_SIZE = 5

type Game struct {
	Players map[string]*state.Player
	Board   gameboard.GameBoard
	CardSet cards.CardSet

	CurrentState state.GameState

	pendingSubmissions map[string][]cards.Card // Player submissions
	pendingSequence    []cards.Card            // Ordered list of all player cards
}

func newGame(players map[string]*state.Player, board gameboard.GameBoard, cardSet cards.CardSet) *Game {

	playerNum := 1
	for _, player := range players {
		player.Deck = cards.NewDeckFromSet(cardSet, len(players), playerNum)
		playerNum += 1
	}

	currentState := state.NewState(0, players, board)

	fmt.Println(fmt.Sprintf("New State: %+v", currentState))

	return &Game{
		Players:            players,
		Board:              board,
		CardSet:            cardSet,
		CurrentState:       DealCards(currentState),
		pendingSubmissions: make(map[string][]cards.Card),
	}
}

func DealCards(inState state.GameState) state.GameState {
	priority := 1
	for i, _ := range inState.Players {
		for len(inState.Players[i].Hand) < HAND_SIZE {
			if len(inState.Players[i].Deck) == 0 {
				inState.Players[i].Deck = cards.ShuffleCards(inState.Players[i].Discard)
				inState.Players[i].Discard = make([]cards.Card, 0)
			}
			drawnCard := inState.Players[i].Deck[0]
			drawnCard.Owner = inState.Players[i].Name
			inState.Players[i].Deck = inState.Players[i].Deck[1:]
			inState.Players[i].Hand = append(inState.Players[i].Hand, drawnCard)
			priority += 1
		}
	}
	return inState
}

func (g *Game) SubmitCards(player string, playerCards []cards.Card) error {
	// TODO validate these cards
	if g.pendingSubmissions[player] != nil {
		return fmt.Errorf("Player already has a pending submission")
	}

	submission := make([]cards.Card, 0)

	for _, c := range playerCards {
		submission = append(submission, c)
	}

	g.pendingSubmissions[player] = submission
	return nil
}

func (g *Game) AggregateTurn() []cards.Card {
	cardOrder := make([]cards.Card, 0)
	for _, pendingCards := range g.pendingSubmissions {
		cardOrder = append(cardOrder, pendingCards...)
	}
	g.pendingSubmissions = make(map[string][]cards.Card)

	sort.Slice(cardOrder, func(i, j int) bool { return cardOrder[i].Priority > cardOrder[j].Priority })
	g.pendingSequence = cardOrder
	return cardOrder
}

func (g *Game) ExecuteTurn() {
	// This should carry out the full step sequence (cards) and calculate all actions that fall out

	// 1. Get starting state
	// g.currentState
	// 2. Execute all cards
	intermState := g.CurrentState
	stepSequence := make([]logic.StepSequence, 0)
	for _, c := range g.pendingSequence {
		// TODO: Game logic here
		fmt.Println(fmt.Sprintf("%+v", c))
		seq, newState := logic.ApplyCard(c, intermState)
		stepSequence = append(stepSequence, seq)
		intermState = newState
	}

	intermState = DealCards(intermState)
	// 3. Save ending state so we can tell players about it
	g.CurrentState = intermState
}
