package logic

import (
	"fmt"

	"github.com/bitDecayGames/LudumDare41/server/cards"
	"github.com/bitDecayGames/LudumDare41/server/state"
	"github.com/bitDecayGames/LudumDare41/server/utils"
)

type Step struct {
	actions []Action
}

type StepSequence struct {
	steps []Step
}

func ApplyCard(c cards.Card, g state.GameState) (StepSequence, state.GameState) {
	// Find and remove card from the player hand
	var affectedPlayer *state.Player
	for i, p := range g.Players {
		if c.Owner == p.Name {
			affectedPlayer = &g.Players[i]
			for i, handCard := range p.Hand {
				if handCard.ID == c.ID {
					p.Hand = append(p.Hand[0:i], p.Hand[i+1:]...)
					p.Discard = append(p.Discard, c)
					break
				}
			}
		}
	}
	switch c.CardType {
	case Card_move_forward_1:
		fmt.Println("MOVE FORWARD")
		affectedPlayer.Pos = utils.VecAdd(affectedPlayer.Pos, affectedPlayer.Facing)
		break
	case Card_move_backward:
		fmt.Println("MOVE BACKWARD")
		break
	}

	return StepSequence{}, g
}
