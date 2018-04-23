package logic

import (
	"strconv"
)

type Action struct {
	ID         string `json:"id"`
	PlayerID   string `json:"playerId"`
	ActionType string `json:"actionType"`
}

var nextActionId = 0

func GetAction(actionType, playerId string) Action {
	val := nextActionId
	nextActionId += 1

	return Action{
		ID:         strconv.Itoa(val),
		PlayerID:   playerId,
		ActionType: actionType,
	}
}

func GetNextActionId() string {
	val := nextActionId
	nextActionId += 1
	return strconv.Itoa(val)
}
