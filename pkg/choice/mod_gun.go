package choice

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/vogtp/eggsis-go/pkg/enemy"
	"github.com/vogtp/eggsis-go/pkg/player"
)

type gun struct {
	*Item
	dmg   int
	reach int32

	lastAttack time.Time
	attackFreq time.Duration
}

func newGun(name string, cost int, dmg int, reach int32, frequency time.Duration, desc string) gun {
	g := gun{
		dmg:        dmg,
		reach:      reach,
		attackFreq: frequency,
	}
	g.Item = &Item{
		Name: name,
		// Image:       "res/egg_strauss.png",
		Description: desc,
		Cost:        cost,
		Modifier: func(p *player.Egg) {

		},
		Action: func(p *player.Egg, enemies []*enemy.Enemy) {
			if time.Since(g.lastAttack) < g.attackFreq {
				return
			}

			r := sdl.Rect{
				X: p.X - g.reach,
				Y: p.Y - g.reach,
				H: p.H + g.reach,
				W: p.W + g.reach,
			}
			for _, e := range enemies {
				if !e.IsDead() && e.HasIntersection(&r) {
					e.LP -= g.dmg - e.Armor
					now := time.Now()
					e.DeathTime = now
					g.lastAttack = now
					return
				}
			}
		},
	}
	return g
}
