package logic

import (
	"github.com/bitDecayGames/LudumDare41/server/cards"
	"github.com/bitDecayGames/LudumDare41/server/gameboard"
	"github.com/bitDecayGames/LudumDare41/server/state"
	"github.com/bitDecayGames/LudumDare41/server/utils"
)

type Step struct {
	actions []Action
}

type StepSequence struct {
	steps []Step
}

func ApplyCard(c cards.Card, g state.GameState) (StepSequence, state.GameState) {
	// Find and remove card from the player hand
	var affectedPlayer *state.Player
	for i, p := range g.Players {
		if c.Owner == p.Name {
			affectedPlayer = &g.Players[i]
			for i, handCard := range affectedPlayer.Hand {
				if handCard.ID == c.ID {
					affectedPlayer.Hand = append(affectedPlayer.Hand[0:i], affectedPlayer.Hand[i+1:]...)
					affectedPlayer.Discard = append(affectedPlayer.Discard, c)
					break
				}
			}
		}
	}

	stepSeq := &StepSequence{}

	switch c.CardType {
	case Card_move_forward_1:
		stepSeq, g = attemptMoveForward(affectedPlayer, stepSeq, g)
	case Card_move_forward_2:
		stepSeq, g = attemptMoveForward(affectedPlayer, stepSeq, g)
		stepSeq, g = attemptMoveForward(affectedPlayer, stepSeq, g)
	case Card_move_forward_3:
		stepSeq, g = attemptMoveForward(affectedPlayer, stepSeq, g)
		stepSeq, g = attemptMoveForward(affectedPlayer, stepSeq, g)
		stepSeq, g = attemptMoveForward(affectedPlayer, stepSeq, g)
	case Card_move_backward_1:
		stepSeq, g = attemptMoveBackwards(affectedPlayer, stepSeq, g)
	case Card_rotate_clockwise:
		stepSeq, g = rotate(affectedPlayer, 90, stepSeq, g)
	}

	return *stepSeq, g
}

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

func attemptMoveForward(player *state.Player, stepSeq *StepSequence, g state.GameState) (*StepSequence, state.GameState) {
	return attemptMove(player, player.Facing, stepSeq, g)
}

func attemptMoveBackwards(player *state.Player, stepSeq *StepSequence, g state.GameState) (*StepSequence, state.GameState) {
	return attemptMove(player, utils.VecScale(player.Facing, -1), stepSeq, g)
}

func attemptMove(player *state.Player, direction utils.Vector, stepSeq *StepSequence, g state.GameState) (*StepSequence, state.GameState) {
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
					actions: []Action{
						otherMove,
						playermove,
					},
				}
				otherPlayer.Pos = utils.VecAdd(otherPlayer.Pos, direction)
				player.Pos = utils.VecAdd(player.Pos, direction)
				stepSeq.steps = append(stepSeq.steps, step)
			}
		} else {
			// free to move
			playermove := FacingToMoveAction(direction, player.Name)
			step := Step{
				actions: []Action{
					playermove,
				},
			}
			player.Pos = utils.VecAdd(player.Pos, direction)
			stepSeq.steps = append(stepSeq.steps, step)
		}
	}
	return stepSeq, g
}

func isEmptyTile(target utils.Vector, g state.GameState) bool {
	if target.X < 0 || target.X >= len(g.Board.Tiles) {
		return false
	}

	if target.Y < 0 || target.Y >= len(g.Board.Tiles[0]) {
		return false
	}

	return g.Board.Tiles[target.X][target.Y].TileType == gameboard.Empty_tile
}

func isPlayerOccupying(target utils.Vector, g state.GameState) bool {
	for _, p := range g.Players {
		if utils.VecEquals(p.Pos, target) {
			return true
		}
	}
	return false
}

func getPlayerAtPos(target utils.Vector, g state.GameState) *state.Player {
	for i, p := range g.Players {
		if utils.VecEquals(p.Pos, target) {
			return &g.Players[i]
		}
	}
	return &state.Player{}
}
