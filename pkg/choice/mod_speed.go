package choice

import "github.com/vogtp/eggsis-go/pkg/player"

func speed(name string, cost int, speed int, desc string) Item {
	return Item{
		Name: name,
		// Image:       "res/egg.png",
		Description: desc,
		Cost:        cost,
		Modifier: func(p *player.Egg) {
			p.Speed =+ p.Speed
		},
	}
}
