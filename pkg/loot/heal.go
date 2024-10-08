package loot

import (
	"log/slog"

	"github.com/vogtp/eggsis-go/pkg/thing"
)

type heal struct {
	a int
}

func (heal) Image() string {
	return "res/heart.png"
}

func (h *heal) Loot(p *thing.Thing) {
	slog.Info("Loot heal", "lp", h.a)
	p.LP += h.a
	//p.MaxLp += h.a
	h.a = 0
}

func Heal(lp int) Loot {
	return &heal{a: lp}
}
