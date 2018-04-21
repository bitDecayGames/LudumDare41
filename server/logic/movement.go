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
	if isEmptyTile(targetPos, g) {
		// check if another player is there
		occupied := isPlayerOccupying(targetPos, g)
		if occupied {
			// Another player is occupying the space we want to go
			otherPlayer := getPlayerAtPos(targetPos, g)
			pushPos := utils.VecAdd(otherPlayer.Pos, direction)
			if isEmptyTile(pushPos, g) && !isPlayerOccupying(pushPos, g) {
				// we push the other player
				otherMove := FacingToMoveAction(direction, otherPlayer.Name)
				playermove := FacingToMoveAction(direction, player.Name)
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
			playermove := FacingToMoveAction(direction, player.Name)
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

func FacingToMoveAction(facing utils.Vector, id string) Action {
	if facing.X == 0 && facing.Y == 1 {
		return MoveNorthAction{PlayerID: id, ActionID: GetNextActionId()}
	} else if facing.X == 0 && facing.Y == -1 {
		return MoveSouthAction{PlayerID: id, ActionID: GetNextActionId()}
	} else if facing.X == 1 && facing.Y == 0 {
		return MoveEastAction{PlayerID: id, ActionID: GetNextActionId()}
	} else if facing.X == -1 && facing.Y == 0 {
		return MoveWestAction{PlayerID: id, ActionID: GetNextActionId()}
	} else {
		panic(fmt.Sprintf("Bad facing received: %v", facing))
	}
}

type MoveNorthAction struct {
	ActionID string
	PlayerID string
}

func (ma MoveNorthAction) GetID() string {
	return ma.ActionID
}

func (ma MoveNorthAction) GetActionType() string {
	return "moveNorthAction"
}

func (ma MoveNorthAction) GetPlayerID() string {
	return ma.PlayerID
}

type MoveSouthAction struct {
	ActionID string
	PlayerID string
}

func (ma MoveSouthAction) GetID() string {
	return ma.ActionID
}

func (ma MoveSouthAction) GetActionType() string {
	return "moveSouthAction"
}

func (ma MoveSouthAction) GetPlayerID() string {
	return ma.PlayerID
}

type MoveEastAction struct {
	ActionID string
	PlayerID string
}

func (ma MoveEastAction) GetID() string {
	return ma.ActionID
}

func (ma MoveEastAction) GetActionType() string {
	return "moveEastAction"
}

func (ma MoveEastAction) GetPlayerID() string {
	return ma.PlayerID
}

type MoveWestAction struct {
	ActionID string
	PlayerID string
}

func (ma MoveWestAction) GetID() string {
	return ma.ActionID
}

func (ma MoveWestAction) GetActionType() string {
	return "moveWestAction"
}

func (ma MoveWestAction) GetPlayerID() string {
	return ma.PlayerID
}