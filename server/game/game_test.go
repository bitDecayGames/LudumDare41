package game

import (
	"fmt"
	"testing"

	"github.com/bitDecayGames/LudumDare41/server/cards"
	"github.com/bitDecayGames/LudumDare41/server/gameboard"
	"github.com/bitDecayGames/LudumDare41/server/lobby"
)

func GetTestGame() *Game {
	gameService := NewGameService()

	lobby, err := lobby.NewLobbyService().NewLobby()
	if err != nil {
		t.Fatal(err)
	}

	lobby.AddPlayer("1")
	lobby.AddPlayer("2")
	board := gameboard.LoadBoard("default")
	cardSet := cards.LoadSet("default")
	game := gameService.NewGame(lobby, board, cardSet)
	return game
}

func TestGameCreation(t *testing.T) {
	gameService := NewGameService()

	lobby := lobby.NewLobbyService().NewLobby()
	lobby.AddPlayer("1")
	lobby.AddPlayer("2")
	board := gameboard.LoadBoard("default")
	cardSet := cards.LoadSet("default")
	g := gameService.NewGame(lobby, board, cardSet)

	if len(g.Players) != 2 {
		t.Error("Lobby players did not carry over into game")
	}
}

func TestCardSubmission(t *testing.T) {
	g := GetTestGame()

	g.DealCards()

	for _, player := range g.Players {
		if len(player.Hand) != HAND_SIZE {
			t.Fatal("Player was not dealt correct hand size")
		}
		fmt.Println(fmt.Sprintf("Player %v cards: %v", player.Name, player.Hand))
		for _, c := range player.Hand {
			if c.Owner != player.Name {
				t.Fatal("Card owner not properly set")
			}
		}

		err := g.SubmitCards(player.Name, player.Hand[0:3])
		if err != nil {
			t.Fatal(err)
		}

		for _, c := range g.pendingSubmissions[player.Name] {
			if player.Name != c.Owner {
				t.Fatal("Card did not have owner assigned properly")
			}
		}

		err = g.SubmitCards(player.Name, player.Hand[0:3])
		if err == nil {
			t.Fatal("Duplicate submission allowed")
		}
	}
}
