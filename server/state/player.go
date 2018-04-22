package state

import (
	"github.com/bitDecayGames/LudumDare41/server/cards"
	"github.com/bitDecayGames/LudumDare41/server/utils"
)

type Player struct {
	Name    string `json:"name"`
	Deck    []cards.Card
	Discard []cards.Card
	Hand    []cards.Card `json:"hand"`

	Pos    utils.Vector `json:"pos"`
	Facing utils.Vector `json:"facing"`
}

func (p *Player) DiscardCard(c cards.Card) {
	for i, handCard := range p.Hand {
		if handCard.ID == c.ID {
			p.Hand = append(p.Hand[0:i], p.Hand[i+1:]...)
			p.Discard = append(p.Discard, c)
			return
		}
	}
}

func (p *Player) DiscardEntireHand() {
	for _, c := range p.Hand {
		p.Discard = append(p.Discard, c)
	}
	p.Hand = make([]cards.Card, 0)
}
