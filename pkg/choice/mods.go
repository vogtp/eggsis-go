package choice

import (
	"log/slog"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/vogtp/eggsis-go/pkg/enemy"
	"github.com/vogtp/eggsis-go/pkg/player"
)

var Mods []Item

func init() {
	Mods = make([]Item, 0)
	Mods = append(Mods, Item{
		Name: "Knife",
		// Image:       "res/egg.png",
		Description: "Do more damage with a knife",
		Cost:        30,
		Modifier: func(e *player.Egg) {
			e.DMG += 10
		},
	})
	Mods = append(Mods, Item{
		Name: "Shield",
		// Image:       "res/egg_eagle.png",
		Description: "Protect your egg",
		Cost:        50,
		Modifier: func(p *player.Egg) {
			p.Armor += 1
		},
	})
	Mods = append(Mods, Item{
		Name: "Health",
		// Image:       "res/egg_strauss.png",
		Description: "More health",
		Cost:        20,
		Modifier: func(p *player.Egg) {
			p.MaxLp += 40
			p.LP = p.MaxLp
		},
	})
	Mods = append(Mods, Item{
		Name: "Wings",
		// Image:       "res/egg_strauss.png",
		Description: "Move faster",
		Cost:        20,
		Modifier: func(p *player.Egg) {
			p.Speed += 4
		},
	})
	Mods = append(Mods, Item{
		Name: "Karate",
		// Image:       "res/egg_strauss.png",
		Description: "Fight faster",
		Cost:        30,
		Modifier: func(p *player.Egg) {
			p.AttackFreq = time.Duration(float64(p.AttackFreq) * 0.8)
			slog.Info("Karate", "attack freq", p.AttackFreq)
		},
	})
	Mods = append(Mods, Item{
		Name: "Gun",
		// Image:       "res/egg_strauss.png",
		Description: "Fight faster",
		Cost:        30,
		Modifier: func(p *player.Egg) {

		},
		Action: func(p *player.Egg, enemies []*enemy.Enemy) {
			dmg := 150
			dist := int32(10)
			r := sdl.Rect{
				X: p.X - dist,
				Y: p.Y - dist,
				H: p.H + dist,
				W: p.W + dist,
			}
			for _, e := range enemies {
				if!e.IsDead() && e.HasIntersection(&r) {
					e.LP -= dmg - e.Armor
					e.DeathTime = time.Now()
					return
				}
			}
		},
	})
}
