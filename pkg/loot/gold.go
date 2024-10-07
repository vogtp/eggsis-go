package loot

import (
	"log/slog"

	"github.com/vogtp/eggsis-go/pkg/player"
)

type gold struct {
	g int
}

func (gold) Image() string {
	return "res/gold.png"
}

func (g *gold) Loot(p *player.Egg) {
	slog.Info("Loot gold", "gold", g)
	p.Gold += g.g
	g.g = 0
}

func Gold(g int) Loot {
	return &gold{g: g}
}
