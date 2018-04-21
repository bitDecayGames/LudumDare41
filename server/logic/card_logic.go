package logic

import (
	"github.com/bitDecayGames/LudumDare41/server/cards"
	"github.com/bitDecayGames/LudumDare41/server/state"
)

type Step struct {
	actions []Action
}

type StepSequence struct {
	steps []Step
}

func ApplyCard(c cards.Card, g state.GameState) (StepSequence, state.GameState) {
	for _, p := range g.Players {
		if c.Owner == p.Name {
			for i, handCard := range p.Hand {
				if handCard.ID == c.ID {
					p.Hand = append(p.Hand[0:i], p.Hand[i+1:]...)
					break
				}
			}
		}
	}
	switch c.CardType {
	case "":
		break
	}
	return StepSequence{}, state.GameState{}
}
