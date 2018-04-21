package game

import (
	"fmt"
	"testing"

	"github.com/bitDecayGames/LudumDare41/server/state"
	"github.com/bitDecayGames/LudumDare41/server/utils"

	"github.com/bitDecayGames/LudumDare41/server/cards"
	"github.com/bitDecayGames/LudumDare41/server/gameboard"
	"github.com/bitDecayGames/LudumDare41/server/lobby"
)

func GetTestGame() (*Game, error) {
	gameService := NewGameService()

	lobby, err := lobby.NewLobbyService().NewLobby()
	if err != nil {
		return nil, err
	}

	lobby.AddPlayer("1")
	lobby.AddPlayer("2")
	board := gameboard.LoadBoard("default")
	cardSet := cards.LoadSet("default")
	game := gameService.NewGame(lobby, board, cardSet)
	return game, nil
}

func TestGameCreation(t *testing.T) {
	gameService := NewGameService()

	lobby, err := lobby.NewLobbyService().NewLobby()
	if err != nil {
		t.Fatal(err)
	}

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
	g, err := GetTestGame()
	if err != nil {
		t.Fatal(err)
	}

	// TODO Make test case for wrong tick scenario
	tick := g.CurrentState.Tick

	for _, player := range g.CurrentState.Players {
		if len(player.Hand) != HAND_SIZE {
			t.Fatal("Player was not dealt correct hand size")
		}
		fmt.Println(fmt.Sprintf("Player %v cards: %v", player.Name, player.Hand))
		for _, c := range player.Hand {
			if c.Owner != player.Name {
				t.Fatal("Card owner not properly set")
			}
		}

		cardIds := []int{}
		for _, card := range player.Hand {
			cardIds = append(cardIds, card.ID)
		}

		err := g.SubmitCards(player.Name, tick, cardIds[0:3])
		if err != nil {
			t.Fatal(err)
		}

		for _, c := range g.pendingSubmissions[player.Name] {
			if player.Name != c.Owner {
				t.Fatal("Card did not have owner assigned properly")
			}
		}

		err = g.SubmitCards(player.Name, tick, cardIds[0:3])
		if err == nil {
			t.Fatal("Duplicate submission allowed")
		}
	}
}

func TestRespawn(t *testing.T) {
	p := state.Player{
		Pos: utils.DeadVector,
	}
	testState := state.GameState{
		Players: []state.Player{p},
		Board: gameboard.GameBoard{
			Tiles: [][]gameboard.Tile{
				{gameboard.Tile{TileType: gameboard.Empty_tile}},
			},
		},
	}

	newState := respawnDeadPlayer(testState)
	if !utils.VecEquals(newState.Players[0].Pos, utils.Vector{X: 0, Y: 0}) {
		t.Fatal("Player not respawned as expected")
	}
}
