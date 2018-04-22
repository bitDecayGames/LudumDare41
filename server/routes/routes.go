package routes

import (
	"github.com/bitDecayGames/LudumDare41/server/game"
	"github.com/bitDecayGames/LudumDare41/server/lobby"
	"github.com/bitDecayGames/LudumDare41/server/pubsub"
	"github.com/gorilla/mux"
)

const (
	apiv1 = "/api/v1"
)

type Routes struct {
	Services     *Services
	pubSubRoutes *PubSubRoutes
	lobbyRoutes  *LobbyRoutes
}

func InitRoutes(r *mux.Router) *Routes {
	services := &Services{
		PubSub: pubsub.NewPubSubService(),
		Lobby:  lobby.NewLobbyService(),
		Game:   game.NewGameService(),
	}

	pubSubRoutes := NewPubSubRoutes(services)
	pubSubRoutes.AddRoutes(r)

	lobbyRoutes := NewLobbyRoutes(services)
	lobbyRoutes.AddRoutes(r)

	return &Routes{
		Services:     services,
		pubSubRoutes: pubSubRoutes,
		lobbyRoutes:  lobbyRoutes,
	}
}
