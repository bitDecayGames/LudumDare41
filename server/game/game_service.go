package game

import (
	"github.com/bitDecayGames/LudumDare41/server/cards"
	"github.com/bitDecayGames/LudumDare41/server/gameboard"
	"github.com/bitDecayGames/LudumDare41/server/lobby"
)

type GameService interface {
	NewGame(*lobby.Lobby, gameboard.GameBoard, cards.CardSet) *Game
}

type Vector struct {
	X int
	Y int
}

type Player struct {
	Name    string
	Deck    []cards.Card
	Discard []cards.Card
	Hand    []cards.Card

	Pos    Vector
	Facing Vector
}

type gameService struct {
	activeGames []Game
}

func NewGameService() GameService {
	return &gameService{}
}

func (gs *gameService) NewGame(lobby *lobby.Lobby, board gameboard.GameBoard, cardSet cards.CardSet) *Game {
	players := make(map[string]*Player)
	for _, player := range lobby.Players {
		players[player] = &Player{
			Name: player,
			Hand: make([]cards.Card, 0),
		}
	}
	return newGame(players, board, cardSet)
}
