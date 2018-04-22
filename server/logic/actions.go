package logic

type Action interface {
	GetID() string
	GetActionType() string
	GetPlayerID() string
}

type MoveAction struct {
	ActionID string
	PlayerID string
}

func (ma MoveAction) GetID() string {
	return ma.ActionID
}

func (ma MoveAction) GetActionType() string {
	return "moveAction"
}

func (ma MoveAction) GetPlayerID() string {
	return ma.PlayerID
}

type TurnAction struct {
	ActionID string
	PlayerID string
}

func (ma TurnAction) GetID() string {
	return ma.ActionID
}

func (ma TurnAction) GetActionType() string {
	return "turnAction"
}

func (ma TurnAction) GetPlayerID() string {
	return ma.PlayerID
}
