package routes

import (
	"log"

	"github.com/bitDecayGames/LudumDare41/server/cards"
	"github.com/bitDecayGames/LudumDare41/server/game"
	"github.com/bitDecayGames/LudumDare41/server/gameboard"
	"github.com/bitDecayGames/LudumDare41/server/lobby"
	"github.com/bitDecayGames/LudumDare41/server/pubsub"
)

type Services struct {
	PubSub pubsub.PubSubService
	Lobby  lobby.LobbyService
	Game   game.GameService
}

func (s *Services) SubmitCards(gameName, playerName string, tick int, cardIds []int) []error {
	game, err := s.Game.GetGame(gameName)
	if err != nil {
		return []error{err}
	}

	err = game.SubmitCards(playerName, tick, cardIds)
	if err != nil {
		return []error{err}
	}

	// Check for advance to next turn
	if game.AreSubmissionsComplete() {
		log.Printf("Starting next turn for game %s at tick %v", game.Name, game.CurrentState.Tick)

		orderedCards := game.AggregateTurn()
		log.Printf("Ordered cards: %+v", orderedCards)
		game.ExecuteTurn()

		log.Printf("Turn complete for game %s at tick %v", game.Name, game.CurrentState.Tick)

		msg := pubsub.Message{
			MessageType: pubsub.GameUpdateMessage,
			ID:          game.Name,
			Tick:        game.CurrentState.Tick,
		}
		return s.PubSub.SendMessage(game.Name, msg)
	}

	return []error{}
}

func (s *Services) CreateGame(lobby *lobby.Lobby) []error {
	// TODO Allow different boards and card sets.
	board := gameboard.LoadBoard("default")
	cardSet := cards.LoadSet("default")
	game := s.Game.NewGame(lobby, board, cardSet)

	// TODO Fix
	// if len(game.Players) < minNumPlayers {
	// 	err := fmt.Errorf("minimum number of %v players not met: %v", minNumPlayers, game.Players)
	// 	return []error{err}
	// }

	// if len(game.Players) > maxNumPlayers {
	// 	err := fmt.Errorf("maximum number of %v players exceeded: %v", maxNumPlayers, game.Players)
	// 	return []error{err}
	// }

	msg := pubsub.Message{
		MessageType: pubsub.GameUpdateMessage,
		ID:          game.Name,
		Tick:        game.CurrentState.Tick,
	}
	return s.PubSub.SendMessage(game.Name, msg)
}
