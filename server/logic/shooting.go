package logic

import (
	"github.com/bitDecayGames/LudumDare41/server/state"
	"github.com/bitDecayGames/LudumDare41/server/utils"
)

func shootMainGun(affectedPlayer *state.Player, stepSeq *StepSequence, g state.GameState) (*StepSequence, state.GameState) {
	// see what direction the player is facing
	found, target := findFirstObstacleInDirection(affectedPlayer, g)

	stepSeq.steps = append(stepSeq.steps,
		Step{
			actions: []Action{
				GetShootAction(affectedPlayer.Name),
			},
		})
	// if player, kill it. If wall, do nothing? (Maybe report what wall was hit?)
	if found && target != nil {
		stepSeq.steps = append(stepSeq.steps,
			Step{
				actions: []Action{
					GetDeathAction(target.Name),
				},
			})
		target.Pos = utils.Vector{X: -1, Y: -1}
	}
	// sequence will always include the player shooting
	return stepSeq, g
}

func findFirstObstacleInDirection(player *state.Player, g state.GameState) (bool, *state.Player) {
	targetPos := player.Pos
	found := false
	for !found {
		targetPos = utils.VecAdd(targetPos, player.Facing)
		if isPlayerOccupying(targetPos, g) {
			// a hit!
			return true, getPlayerAtPos(targetPos, g)
		}
		if isEmptyTile(targetPos, g) {
			// round still can travel
			continue
		} else {
			// hit a wall
			return false, nil
		}
	}
	// shouldn't get here. This means that we never found a player or a wall
	return false, nil
}

func GetShootAction(ID string) Action {
	return ShootMainGunsAction{PlayerID: ID, ActionID: GetNextActionId()}
}

type ShootMainGunsAction struct {
	ActionID string
	PlayerID string
}

func (smga ShootMainGunsAction) GetID() string {
	return smga.ActionID
}

func (smga ShootMainGunsAction) GetActionType() string {
	return Action_shoot_main_gun
}

func (smga ShootMainGunsAction) GetPlayerID() string {
	return smga.PlayerID
}

func GetDeathAction(ID string) Action {
	return DeathAction{PlayerID: ID, ActionID: GetNextActionId()}
}

type DeathAction struct {
	ActionID string
	PlayerID string
}

func (da DeathAction) GetID() string {
	return da.ActionID
}

func (da DeathAction) GetActionType() string {
	return Action_death
}

func (da DeathAction) GetPlayerID() string {
	return da.PlayerID
}
