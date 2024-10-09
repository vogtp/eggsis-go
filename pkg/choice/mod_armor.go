package choice

import "github.com/vogtp/eggsis-go/pkg/player"

func damage(name string, cost int, armor int, desc string) Item {
	return Item{
		Name: name,
		// Image:       "res/egg.png",
		Description: desc,
		Cost:        cost,
		Modifier: func(p *player.Egg) {
			p.Armor += armor
		},
	}
}
