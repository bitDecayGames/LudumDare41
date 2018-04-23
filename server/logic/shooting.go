package logic

import (
	"github.com/bitDecayGames/LudumDare41/server/state"
	"github.com/bitDecayGames/LudumDare41/server/utils"
)

func shootMainGun(affectedPlayer *state.Player, stepSeq []Step, g state.GameState) ([]Step, state.GameState) {
	// see what direction the player is facing
	found, target := findFirstObstacleInDirection(affectedPlayer, g)

	stepSeq = append(stepSeq,
		Step{
			Actions: []Action{
				GetAction(Action_shoot_main_gun, affectedPlayer.Name),
			},
		})
	// if player, kill it. If wall, do nothing? (Maybe report what wall was hit?)
	if found && target != nil {
		stepSeq = append(stepSeq,
			Step{
				Actions: []Action{
					GetAction(Action_death, target.Name),
				},
			})
		target.DiscardEntireHand()
		target.Pos = utils.Vector{X: -1, Y: -1}
	}
	// sequence will always include the player shooting
	return stepSeq, g
}

func findFirstObstacleInDirection(player *state.Player, g state.GameState) (bool, *state.Player) {
	targetPos := player.Pos
	found := false
	for !found {
		targetPos = utils.VecAdd(targetPos, player.Facing)
		if isPlayerOccupying(targetPos, g) {
			// a hit!
			return true, getPlayerAtPos(targetPos, g)
		}
		if isEmptyTile(targetPos, g) {
			// round still can travel
			continue
		} else {
			// hit a wall
			return false, nil
		}
	}
	// shouldn't get here. This means that we never found a player or a wall
	return false, nil
}
