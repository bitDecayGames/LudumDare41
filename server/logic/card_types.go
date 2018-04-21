package logic

import (
	"github.com/bitDecayGames/LudumDare41/server/cards"
)

const Card_move_forward_1 = "moveForward1Card"
const Card_move_backward = "moveBackwardCard"

var CardTypeMap = map[string]cards.Card{
	Card_move_forward_1: cards.Card{
		ID:       -1,
		CardType: Card_move_forward_1,
	},
}
