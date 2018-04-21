package game

import (
	"testing"

	"github.com/bitDecayGames/LudumDare41/server/cards"
	"github.com/bitDecayGames/LudumDare41/server/gameboard"
	"github.com/bitDecayGames/LudumDare41/server/lobby"
)

func TestGameCreation(t *testing.T) {
	gameService := NewGameService()

	lobby := lobby.NewLobbyService().NewLobby()
	lobby.AddPlayer("1")
	lobby.AddPlayer("2")
	board := gameboard.NewBoard("default")
	cardSet := cards.LoadSet("default")
	game := gameService.NewGame(lobby, board, cardSet)

	if len(game.Players) != 2 {
		t.Error("Lobby players did not carry over into game")
	}

	if game.Board == nil {
		t.Error("Game board is nil")
	}
}
