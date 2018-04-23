package logic

import (
	"fmt"

	"github.com/bitDecayGames/LudumDare41/server/state"
	"github.com/bitDecayGames/LudumDare41/server/utils"
)

func attemptMoveForward(player *state.Player, stepSeq []Step, g state.GameState) ([]Step, state.GameState) {
	return attemptMove(player, player.Facing, stepSeq, g)
}

func attemptMoveBackwards(player *state.Player, stepSeq []Step, g state.GameState) ([]Step, state.GameState) {
	return attemptMove(player, utils.VecScale(player.Facing, -1), stepSeq, g)
}

// TODO: after a movement, we need to check for collecting a crate
func attemptMove(player *state.Player, direction utils.Vector, stepSeq []Step, g state.GameState) ([]Step, state.GameState) {
	targetPos := utils.VecAdd(player.Pos, direction)
	if IsEmptyTile(targetPos, g) {
		// check if another player is there
		occupied := IsPlayerOccupying(targetPos, g)
		if occupied {
			// Another player is occupying the space we want to go
			otherPlayer := GetPlayerAtPos(targetPos, g)
			pushPos := utils.VecAdd(otherPlayer.Pos, direction)
			if IsEmptyTile(pushPos, g) && !IsPlayerOccupying(pushPos, g) {
				// we push the other player
				otherMove := FacingToMoveAction(direction, otherPlayer)
				playermove := FacingToMoveAction(direction, player)
				step := Step{
					Actions: []Action{
						otherMove,
						playermove,
					},
				}
				otherPlayer.Pos = utils.VecAdd(otherPlayer.Pos, direction)
				player.Pos = utils.VecAdd(player.Pos, direction)
				stepSeq = append(stepSeq, step)
			}
		} else {
			// free to move
			playermove := FacingToMoveAction(direction, player)
			step := Step{
				Actions: []Action{
					playermove,
				},
			}
			player.Pos = utils.VecAdd(player.Pos, direction)
			stepSeq = append(stepSeq, step)
		}
	}
	return stepSeq, g
}

func FacingToMoveAction(facing utils.Vector, p *state.Player) Action {
	destPos := utils.VecAdd(p.Pos, facing)
	if facing.X == 0 && facing.Y == 1 {
		return GetAction(Action_move_north, p.Name, destPos)
	} else if facing.X == 0 && facing.Y == -1 {
		return GetAction(Action_move_south, p.Name, destPos)
	} else if facing.X == 1 && facing.Y == 0 {
		return GetAction(Action_move_east, p.Name, destPos)
	} else if facing.X == -1 && facing.Y == 0 {
		return GetAction(Action_move_west, p.Name, destPos)
	} else {
		panic(fmt.Sprintf("Bad facing received: %v", facing))
	}
}
