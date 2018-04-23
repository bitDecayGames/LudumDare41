package logic

import (
	"github.com/bitDecayGames/LudumDare41/server/utils"

	"github.com/bitDecayGames/LudumDare41/server/cards"
	"github.com/bitDecayGames/LudumDare41/server/state"
)

type Step struct {
	Actions []Action `json:"actions"`
}

type StepSequence struct {
	Cards []cards.Card `json:"cards"`
	Steps []Step       `json:"steps"`
}

func ApplyCard(c cards.Card, g state.GameState) ([]Step, state.GameState) {
	steps := make([]Step, 0)
	// Find and remove card from the player hand

	var affectedPlayer *state.Player
	for i, p := range g.Players {
		if c.Owner == p.Name {
			affectedPlayer = &g.Players[i]

			if utils.VecEquals(affectedPlayer.Pos, utils.DeadVector) {
				// player is dead, don't play any more of their cards
				steps = append(steps, Step{
					Actions: []Action{
						GetAction(Action_dispose_next_card, affectedPlayer.Name, affectedPlayer.Pos),
					},
				})
				return steps, g
			}

			steps = append(steps, Step{
				Actions: []Action{
					GetAction(Action_play_next_card, affectedPlayer.Name, affectedPlayer.Pos),
				},
			})
			affectedPlayer.DiscardCard(c)
		}
	}

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

	// check for map interaction (crate pickups, etc)
	for i, p := range g.Players {
		if !utils.VecEquals(p.Pos, utils.DeadVector) && utils.VecEquals(p.Pos, g.Crate) {
			g.Players[i].Points += 1
			steps = append(steps, Step{
				Actions: []Action{
					GetAction(Action_collect_crate, p.Name, g.Crate),
				},
			})
			g.Crate = utils.DeadVector

			if g.Players[i].Points >= POINTS_FOR_VICTORY {
				steps = append(steps, Step{
					Actions: []Action{
						GetAction(Action_victory, p.Name, p.Pos),
					},
				})
			}
		}
	}

	// set our player data back into the game state
	// for i, p := range g.Players {
	// 	if p.Name == affectedPlayer.Name {
	// 		g.Players[i] = *affectedPlayer
	// 	}
	// }

	return steps, g
}
