package player

import "github.com/vogtp/eggsis-go/pkg/enemy"

type ActionFunc func(*Egg, []*enemy.Enemy)

func (p *Egg) Action(enemies []*enemy.Enemy) {
	for _, a := range p.Actions {
		if a == nil {
			continue
		}
		a(p, enemies)
	}
}

