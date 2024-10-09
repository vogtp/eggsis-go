package choice

import (
	"time"

	"github.com/vogtp/eggsis-go/pkg/player"
)

var Players []Item

func init() {
	Players = make([]Item, 0)
	normalEgg := Item{
		Name:        "Egg",
		Image:       "res/egg.png",
		Description: "Just a plain old egg!",
		Modifier:    func(p *player.Egg) {
		p.Gold+=50	
		},
	
	}
	Players = append(Players, normalEgg)
	Players = append(Players, Item{
		Name:        "Eagle Egg",
		Image:       "res/egg_eagle.png",
		Description: "An attackin egg...",
		Modifier: func(p *player.Egg) {
			p.DMG+=7
			p.Speed+=3
		},
	})
	Players = append(Players, Item{
		Name:        "Staussen Egg",
		Image:       "res/egg_strauss.png",
		Description: "An defensive egg...",
		Modifier: func(p *player.Egg) {
			p.Armor+=1
			p.MaxLp+=30
			p.LP = p.MaxLp
		},
	})
	Players = append(Players, Item{
		Name:        "Egg Stinguish",
		Image:       "res/atomic.png",
		Description: "Completely OP",
		Modifier: func(p *player.Egg) {
			p.Armor=20
			p.DMG=20
			p.AttackFreq=10*time.Millisecond
		},
	})
	Players = append(Players, Item{
		Name:        "Egg Scuse Me",
		Image:       "res/egg_excuse.png",
		Description: "Sorry about it...",
		Modifier: func(p *player.Egg) {
			p.Armor=200
			p.DMG=0		},
	})
}
