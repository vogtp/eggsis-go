package enemy

import "github.com/vogtp/eggsis-go/pkg/player"

func calcDmg(dmg int, armor int) int {
	d := dmg - armor
	if d < 0 {
		d = 0
	}
	return d
}

func (e *Enemy) Fight(p *player.Player) {
	p.LP -= calcDmg(e.DMG, p.Armor)
	e.LP -= calcDmg(p.DMG, e.Armor)
}
