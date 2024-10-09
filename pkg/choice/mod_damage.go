package choice

import "github.com/vogtp/eggsis-go/pkg/player"

func armor(name string, cost int, dmg int, desc string) Item {
	return Item{
		Name: name,
		// Image:       "res/egg.png",
		Description: desc,
		Cost:        cost,
		Modifier: func(p *player.Egg) {
			p.DMG += dmg
		},
	}
}
