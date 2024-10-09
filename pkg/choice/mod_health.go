package choice

import "github.com/vogtp/eggsis-go/pkg/player"

func health(name string, cost int, lp int, desc string) Item {
	return Item{
		Name: name,
		// Image:       "res/egg.png",
		Description: desc,
		Cost:        cost,
		Modifier: func(p *player.Egg) {
			p.MaxLp =+ lp
			p.LP = p.MaxLp
		},
	}
}
