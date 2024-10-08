package enemy

import (
	"log/slog"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/vogtp/eggsis-go/pkg/player"
	"github.com/vogtp/eggsis-go/pkg/thing"
)

func calcDmg(attacker *thing.Thing, target *thing.Thing) int {
	if time.Since(attacker.LastAttack) < attacker.AttackFreq {
		return 0
	}
	attacker.LastAttack = time.Now()
	d := attacker.DMG - target.Armor
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
	p.LP -= calcDmg(e.Thing, p.Thing)
	e.LP -= calcDmg(p.Thing, e.Thing)
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
