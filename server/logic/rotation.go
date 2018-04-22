package logic

import "github.com/bitDecayGames/LudumDare41/server/state"

func rotate(player *state.Player, degrees int, stepSeq *StepSequence, g state.GameState) (*StepSequence, state.GameState) {
	var newX int
	var newY int

	switch degrees {
	case 90:
		newX = player.Facing.Y
		newY = player.Facing.X * -1
	case -90:
		newX = player.Facing.Y * -1
		newY = player.Facing.X
	case 180:
		newX = player.Facing.X * -1
		newY = player.Facing.Y * -1
	}

	player.Facing.X = newX
	player.Facing.Y = newY

	step := Step{
		actions: []Action{
			DegreesToRotateAction(degrees, player.Name),
		},
	}
	stepSeq.steps = append(stepSeq.steps, step)

	return stepSeq, g
}

func DegreesToRotateAction(degrees int, id string) Action {
	switch degrees {
	case 90:
		return RotateCounterClockwiseAction{PlayerID: id, ActionID: GetNextActionId()}
	case -90:
		return RotateClockwiseAction{PlayerID: id, ActionID: GetNextActionId()}
	case 180:
		fallthrough
	default:
		return Rotate180Action{PlayerID: id, ActionID: GetNextActionId()}
	}
}

type RotateCounterClockwiseAction struct {
	ActionID string
	PlayerID string
}

func (rcca RotateCounterClockwiseAction) GetID() string {
	return rcca.ActionID
}

func (rcca RotateCounterClockwiseAction) GetActionType() string {
	return Action_rotate_counter_clockwise
}

func (rcca RotateCounterClockwiseAction) GetPlayerID() string {
	return rcca.PlayerID
}

type RotateClockwiseAction struct {
	ActionID string
	PlayerID string
}

func (rca RotateClockwiseAction) GetID() string {
	return rca.ActionID
}

func (rca RotateClockwiseAction) GetActionType() string {
	return Action_rotate_clockwise
}

func (rca RotateClockwiseAction) GetPlayerID() string {
	return rca.PlayerID
}

type Rotate180Action struct {
	ActionID string
	PlayerID string
}

func (r1a Rotate180Action) GetID() string {
	return r1a.ActionID
}

func (r1a Rotate180Action) GetActionType() string {
	return Action_rotate_180
}

func (r1a Rotate180Action) GetPlayerID() string {
	return r1a.PlayerID
}
