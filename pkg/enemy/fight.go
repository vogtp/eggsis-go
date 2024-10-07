package enemy

import (
	"log/slog"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/vogtp/eggsis-go/pkg/player"
)

func calcDmg(dmg int, armor int) int {
	d := dmg - armor
	if d < 0 {
		d = 0
	}
	return d
}

func (e *Enemy) Fight(p *player.Egg) {
	slog.Debug("fight", "enemy", e, "player", p)
	if e.IsDead() {
		e.GetLooted(p)
		return
	}
	p.LP -= calcDmg(e.DMG, p.Armor)
	e.LP -= calcDmg(p.DMG, e.Armor)
	if e.IsDead() {
		e.DeathTime = time.Now()
	}
}

func (e *Enemy) IsDead() bool {
	d := e.Thing.IsDead()

	if d && e.LP != -255 {
		e.LP = -255

		if suf, err := img.Load(e.LootDrop.Image()); err == nil {
			suf.SetAlphaMod(250)
			e.Surface = suf
		} else {
			slog.Warn("cannot load dead img", "err", err)
		}
	}
	return d
}
