package loot

import (
	"log/slog"

	"github.com/vogtp/eggsis-go/pkg/player"
)

func Gold(g int) Loot {
	return func(p *player.Player) {
		slog.Info("Loot gold", "gold", g)
		p.Gold += g
	}
}
