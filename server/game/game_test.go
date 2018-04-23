package game

import (
	"fmt"
	"testing"

	"github.com/bitDecayGames/LudumDare41/server/logic"

	"github.com/bitDecayGames/LudumDare41/server/cards"
	"github.com/bitDecayGames/LudumDare41/server/gameboard"
	"github.com/bitDecayGames/LudumDare41/server/lobby"
	"github.com/bitDecayGames/LudumDare41/server/state"
	"github.com/bitDecayGames/LudumDare41/server/utils"
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

	step, newState := respawnDeadPlayer(testState)
	if !utils.VecEquals(newState.Players[0].Pos, utils.Vector{X: 0, Y: 0}) {
		t.Fatal("Player not respawned as expected")
	}

	if len(step.Actions) != 1 || step.Actions[0].GetActionType() != logic.Action_spawn {
		t.Errorf("Spawn action didn't return, got: %+v", step)
		t.Errorf("Length of actions: %v", len(step.Actions))
		t.Fatalf("Action Type %v", step.Actions[0].GetActionType())
	}
}

func TestCardOrdering(t *testing.T) {
	g := Game{
		pendingSubmissions: make(map[string][]cards.Card),
	}

	g.pendingSubmissions["one"] = []cards.Card{
		cards.Card{ID: 5, Priority: 5},
		cards.Card{ID: 1, Priority: 1},
		cards.Card{ID: 3, Priority: 3},
	}

	g.pendingSubmissions["two"] = []cards.Card{
		cards.Card{ID: 2, Priority: 2},
		cards.Card{ID: 9, Priority: 9},
		cards.Card{ID: 4, Priority: 4},
	}

	order := g.AggregateTurn()
	if order[0].ID != 5 ||
		order[1].ID != 2 ||
		order[2].ID != 9 ||
		order[3].ID != 1 ||
		order[4].ID != 4 ||
		order[5].ID != 3 {
		t.Fatal("Card turn order was not correct")
	}
}
