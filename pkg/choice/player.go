package choice

import "github.com/vogtp/eggsis-go/pkg/player"

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
}
