package choice

import (
	"log/slog"
	"time"

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
	Mods = append(Mods, *newGun("Old gun", 20, 3, 20, 500*time.Millisecond, "Old slow gun").Item)
	Mods = append(Mods, *newGun("Good gun", 80, 10, 100, 100*time.Millisecond, "Strong fast gun gun").Item)
	Mods = append(Mods, *newGun("Machine gun", 200, 30, 120, 50*time.Millisecond, "Machine gun").Item)
}
