package state

import (
	"github.com/bitDecayGames/LudumDare41/server/cards"
	"github.com/bitDecayGames/LudumDare41/server/utils"
)

type Player struct {
	Name    string ``
	Deck    []cards.Card
	Discard []cards.Card
	Hand    []cards.Card ``

	Pos    utils.Vector ``
	Facing utils.Vector ``
}
