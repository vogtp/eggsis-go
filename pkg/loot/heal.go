package loot

import (
	"log/slog"

	"github.com/vogtp/eggsis-go/pkg/player"
)

func Heal(lp int) Loot {
	return func(p *player.Player) {
		slog.Info("Loot heal", "lp", lp)
		p.LP += lp
	}
}
