package logic

import (
	"fmt"

	"github.com/bitDecayGames/LudumDare41/server/cards"
	"github.com/bitDecayGames/LudumDare41/server/state"
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
			for i, handCard := range affectedPlayer.Hand {
				if handCard.ID == c.ID {
					affectedPlayer.Hand = append(affectedPlayer.Hand[0:i], affectedPlayer.Hand[i+1:]...)
					affectedPlayer.Discard = append(affectedPlayer.Discard, c)
					break
				}
			}
		}
	}
	if affectedPlayer == nil {
		fmt.Println("THE PLAYER WAS MUFFUGGIN NIL")
	}

	stepSeq := &StepSequence{}

	switch c.CardType {
	case Card_move_forward_1:
		stepSeq, g = attemptMoveForward(affectedPlayer, stepSeq, g)
	case Card_move_forward_2:
		stepSeq, g = attemptMoveForward(affectedPlayer, stepSeq, g)
		stepSeq, g = attemptMoveForward(affectedPlayer, stepSeq, g)
	case Card_move_forward_3:
		stepSeq, g = attemptMoveForward(affectedPlayer, stepSeq, g)
		stepSeq, g = attemptMoveForward(affectedPlayer, stepSeq, g)
		stepSeq, g = attemptMoveForward(affectedPlayer, stepSeq, g)
	case Card_move_backward_1:
		stepSeq, g = attemptMoveBackwards(affectedPlayer, stepSeq, g)
	case Card_rotate_clockwise:
		stepSeq, g = rotate(affectedPlayer, 90, stepSeq, g)
	case Card_shoot_main_turret:
		stepSeq, g = shootMainGun(affectedPlayer, stepSeq, g)
	}

	return *stepSeq, g
}
