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
	testRoutes   *TestRoutes
	pubSubRoutes *PubSubRoutes
	lobbyRoutes  *LobbyRoutes
	gameRoutes   *GameRoutes
}

func InitRoutes(r *mux.Router) *Routes {
	services := &Services{
		PubSub: pubsub.NewPubSubService(),
		Lobby:  lobby.NewLobbyService(),
		Game:   game.NewGameService(),
	}

	testRoutes := NewTestRoutes(services)
	testRoutes.AddRoutes(r)

	pubSubRoutes := NewPubSubRoutes(services)
	pubSubRoutes.AddRoutes(r)

	lobbyRoutes := NewLobbyRoutes(services)
	lobbyRoutes.AddRoutes(r)

	gameRoutes := NewGameRoutes(services)
	gameRoutes.AddRoutes(r)

	return &Routes{
		Services:     services,
		pubSubRoutes: pubSubRoutes,
		lobbyRoutes:  lobbyRoutes,
	}
}
