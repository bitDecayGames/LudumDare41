package test_scripts

import (
	"fmt"
	"math"
	"testing"

	"github.com/bitDecayGames/LudumDare41/server/cards"
	"github.com/bitDecayGames/LudumDare41/server/gameboard"

	"github.com/bitDecayGames/LudumDare41/server/game"
	"github.com/bitDecayGames/LudumDare41/server/lobby"
)

func TestFullRun(t *testing.T) {
	lobbyServer := lobby.NewLobbyService()
	lobby := lobbyServer.NewLobby()

	lobby.AddPlayer("Logan")
	lobby.AddPlayer("Jake")

	board := gameboard.LoadBoard("default")
	cardSet := cards.LoadSet("default")

	gameServer := game.NewGameService()
	g := gameServer.NewGame(lobby, board, cardSet)

	if len(g.Players) < 2 {
		t.Fatal("Not all players are in game")
	}

	g.DealCards()

	for _, player := range g.Players {
		if len(player.Hand) != game.HAND_SIZE {
			t.Fatal("Player was not dealt correct hand size")
		}
		fmt.Println(fmt.Sprintf("Player %v cards: %v", player.Name, player.Hand))
		err := g.SubmitCards(player.Name, player.Hand[0:3])
		if err != nil {
			t.Fatal(err)
		}

		err = g.SubmitCards(player.Name, player.Hand[0:3])
		if err == nil {
			t.Fatal("Duplicate submission allowed")
		}
	}

	turnCards := g.AggregateTurn()
	fmt.Println(turnCards)
	lastValue := math.MaxInt64
	for _, card := range turnCards {
		if card.Priority >= lastValue {
			t.Fatal("Cards are not properly priority sorted")
		}
		lastValue = card.Priority
	}
}
