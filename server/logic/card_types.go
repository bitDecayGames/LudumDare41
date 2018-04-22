package logic

import (
	"github.com/bitDecayGames/LudumDare41/server/cards"
)

const Card_move_forward_1 = "moveForward1Card"
const Card_move_forward_2 = "moveForward2Card"
const Card_move_forward_3 = "moveForward3Card"

const Card_move_backward_1 = "moveBackwardCard"

var CardTypeMap = map[string]cards.Card{
	Card_move_forward_1: cards.Card{
		ID:       -1,
		CardType: Card_move_forward_1,
	},
	Card_move_forward_2: cards.Card{
		ID:       -1,
		CardType: Card_move_forward_2,
	},
	Card_move_forward_3: cards.Card{
		ID:       -1,
		CardType: Card_move_forward_3,
	},

	Card_move_backward_1: cards.Card{
		ID:       -1,
		CardType: Card_move_backward_1,
	},
}
