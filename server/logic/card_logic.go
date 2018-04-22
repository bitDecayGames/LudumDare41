package logic

import (
	"fmt"

	"github.com/bitDecayGames/LudumDare41/server/cards"
	"github.com/bitDecayGames/LudumDare41/server/state"
)

type Step struct {
	Actions []Action
}

type StepSequence struct {
	Steps []Step
}

func ApplyCard(c cards.Card, g state.GameState) ([]Step, state.GameState) {
	// Find and remove card from the player hand
	var affectedPlayer *state.Player
	for i, p := range g.Players {
		if c.Owner == p.Name {
			affectedPlayer = &g.Players[i]
			affectedPlayer.DiscardCard(c)
		}
	}
	if affectedPlayer == nil {
		fmt.Println("THE PLAYER WAS MUFFUGGIN NIL")
	}

	steps := make([]Step, 0)

	switch c.CardType {
	case Card_move_forward_1:
		steps, g = attemptMoveForward(affectedPlayer, steps, g)
	case Card_move_forward_2:
		steps, g = attemptMoveForward(affectedPlayer, steps, g)
		steps, g = attemptMoveForward(affectedPlayer, steps, g)
	case Card_move_forward_3:
		steps, g = attemptMoveForward(affectedPlayer, steps, g)
		steps, g = attemptMoveForward(affectedPlayer, steps, g)
		steps, g = attemptMoveForward(affectedPlayer, steps, g)
	case Card_move_backward_1:
		steps, g = attemptMoveBackwards(affectedPlayer, steps, g)
	case Card_rotate_clockwise:
		steps, g = rotate(affectedPlayer, 90, steps, g)
	case Card_shoot_main_turret:
		steps, g = shootMainGun(affectedPlayer, steps, g)
	}

	for i, p := range g.Players {
		if p.Name == affectedPlayer.Name {
			g.Players[i] = *affectedPlayer
		}
	}

	return steps, g
}
