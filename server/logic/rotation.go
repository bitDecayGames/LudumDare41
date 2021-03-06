package logic

import "github.com/bitDecayGames/LudumDare41/server/state"

func rotate(player *state.Player, degrees int, stepSeq []Step, g state.GameState) ([]Step, state.GameState) {
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
		Actions: []Action{
			DegreesToRotateAction(degrees, player),
		},
	}
	stepSeq = append(stepSeq, step)

	return stepSeq, g
}

func DegreesToRotateAction(degrees int, p *state.Player) Action {
	switch degrees {
	case 90:
		return GetAction(Action_rotate_clockwise, p.Name, p.Pos)
	case -90:
		return GetAction(Action_rotate_counter_clockwise, p.Name, p.Pos)
	case 180:
		fallthrough
	default:
		return GetAction(Action_rotate_180, p.Name, p.Pos)
	}
}
