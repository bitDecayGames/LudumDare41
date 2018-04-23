package logic

import (
	"strconv"

	"github.com/bitDecayGames/LudumDare41/server/utils"
)

type Action struct {
	ID         string       `json:"id"`
	PlayerID   string       `json:"playerId"`
	ActionType string       `json:"actionType"`
	Pos        utils.Vector `json:"position"`
}

var nextActionId = 0

func GetAction(actionType, playerId string, where utils.Vector) Action {
	val := nextActionId
	nextActionId += 1

	return Action{
		ID:         strconv.Itoa(val),
		PlayerID:   playerId,
		ActionType: actionType,
		Pos:        where,
	}
}

func GetNextActionId() string {
	val := nextActionId
	nextActionId += 1
	return strconv.Itoa(val)
}
