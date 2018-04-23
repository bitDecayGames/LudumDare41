package test_scripts

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/bitDecayGames/LudumDare41/server/logic"

	"github.com/bitDecayGames/LudumDare41/server/cards"
	"github.com/bitDecayGames/LudumDare41/server/gameboard"

	"github.com/bitDecayGames/LudumDare41/server/game"
	"github.com/bitDecayGames/LudumDare41/server/lobby"
)

func TestFullRun(t *testing.T) {
	lobbyServer := lobby.NewLobbyService()
	lobby, err := lobbyServer.NewLobby()
	if err != nil {
		t.Fatal(err)
	}

	lobby.AddPlayer("Logan")
	lobby.AddPlayer("Jake")

	board := gameboard.LoadBoard("default")
	cardSet := cards.LoadSet(logic.CardSetMap["debug"])

	gameServer := game.NewGameService()
	g := gameServer.NewGame(lobby, board, cardSet)

	if len(g.Players) < 2 {
		t.Fatal("Not all players are in game")
	}

	g.AggregateTurn()
	g.ExecuteTurn()

	for _, player := range g.CurrentState.Players {
		fmt.Println(fmt.Sprintf("Player %v cards: %v", player.Name, player.Hand))

		cardIds := []int{}
		for _, card := range player.Hand {
			cardIds = append(cardIds, card.ID)
		}

		err := g.SubmitCards(player.Name, g.CurrentState.Tick, cardIds[0:3])
		if err != nil {
			t.Fatal(err)
		}
	}

	g.AggregateTurn()
	g.ExecuteTurn()

	bytes, err := json.Marshal(g.LastSequence)
	log.Println(err)
	fmt.Println(fmt.Sprintf("Dat Json: %v", string(bytes)))

	for _, p := range g.CurrentState.Players {
		if len(p.Hand) != 5 {
			t.Fatal("Cards were not replenished")
		}
	}
}
