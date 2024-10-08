package choice

import "github.com/vogtp/eggsis-go/pkg/player"

var Mods []Item

func init() {
	Mods = make([]Item, 0)
	Mods = append(Mods, Item{
		Name:        "Knife",
		// Image:       "res/egg.png",
		Description: "Do more damage with a knife",
		Cost: 50,
		Modifier:    func(e *player.Egg) {
			e.DMG +=10
		},
	})
	Mods = append(Mods, Item{
		Name:        "Shield",
		// Image:       "res/egg_eagle.png",
		Description: "Protect your egg",
		Cost: 50,
		Modifier: func(p *player.Egg) {
			p.Armor += 3
		},
	})
	Mods = append(Mods, Item{
		Name:        "Health",
		// Image:       "res/egg_strauss.png",
		Description: "More health",
		Cost: 50,
		Modifier: func(p *player.Egg) {
			p.MaxLp += 30
			p.LP = p.MaxLp
		},
	})
	Mods = append(Mods, Item{
		Name:        "Wings",
		// Image:       "res/egg_strauss.png",
		Description: "Move faster",
		Cost: 50,
		Modifier: func(p *player.Egg) {
			p.Speed += 3
		},
	})
}
