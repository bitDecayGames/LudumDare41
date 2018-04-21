package main

import "github.com/bitDecayGames/LudumDare41/server/lobby"

const (
	jacque = "Jacque"
	monday = "Monday"
)

func main() {
	// create 2 player game
	lobbyService := lobby.NewLobbyService()
	lobby := lobbyService.NewLobby()

	lobby.AddPlayer(jacque)
	lobby.AddPlayer(monday)

	// deck := deck.LoadDeck("default")

	// map := map.LoadMap("testLevel")
	// gameInstance := game.NewGame(lobby, map, deck)

	// gameInstance.Start()

	// gameInstance.DealCards()
	// gameInstance.SendUpdates() // This sends SanitizedGameStates to each player to set initial state

	// gameInstance.SubmitCards(jacque, cards)
	// gameInstance.SubmitCards(monday, cards)

	// gameInstance.PrioritizeCards()

	// gameInstance.ResolveCards()

	// if gameInstance.GameOver() {
	// 	os.Exit(1)
	// }

	// gameInstance.DealCards()
	// gameInstance.SendUpdates()
}

// Card represents a single unique playing card
type Card struct {
	ID int
}

// GameState is a fully encompassing snapshot of the state of the game
type GameState struct {
}

// SanitizedGameState is a copy of the game state tailored to a single player. It only contains information pertinent to that player
type SanitizedGameState struct {
}

// TurnSubmission represents a given player's card choices for the given turn
type TurnSubmission struct {
	PlayerID int
	Tick     int
	Cards    []Card
}

// TickUpdate is a data update intended to be sent to the client that contains all needed information on the start state and end state of a given game tick
type TickUpdate struct {
	Tick       int
	PlayerID   int
	StartState GameState
	EndState   GameState
	// Sequences
}
