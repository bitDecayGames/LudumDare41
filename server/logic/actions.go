package logic

import (
	"fmt"
	"strconv"

	"github.com/bitDecayGames/LudumDare41/server/utils"
)

var nextActionId = 0

func GetNextActionId() string {
	val := nextActionId
	nextActionId += 1
	return strconv.Itoa(val)
}

type Action interface {
	GetID() string
	GetActionType() string
	GetPlayerID() string
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
