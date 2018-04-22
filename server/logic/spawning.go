package logic

func GetSpawnAction(id string) Action {
	return SpawnAction{PlayerID: id, ActionID: GetNextActionId()}
}

type SpawnAction struct {
	ActionID string
	PlayerID string
}

func (sa SpawnAction) GetID() string {
	return sa.ActionID
}

func (ra SpawnAction) GetActionType() string {
	return Action_spawn
}

func (sa SpawnAction) GetPlayerID() string {
	return sa.PlayerID
}
