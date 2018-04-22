package logic

import (
	"strconv"
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
