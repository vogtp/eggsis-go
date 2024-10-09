package choice

import (
	"time"

	"github.com/vogtp/eggsis-go/pkg/player"
)

func attackSpeed(name string, cost int, freq float64, desc string) Item {
	return Item{
		Name: name,
		// Image:       "res/egg.png",
		Description: desc,
		Cost:        cost,
		Modifier: func(p *player.Egg) {
			p.AttackFreq = time.Duration(float64(p.AttackFreq) * freq)
		},
	}
}
