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
