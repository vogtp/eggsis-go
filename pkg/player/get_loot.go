package player

import (
	"log/slog"
	"time"

	"github.com/vogtp/eggsis-go/pkg/enemy"
)

func (p *Egg) GetLooted(e *enemy.Enemy) {
	if e.LootDrop == nil {
		slog.Debug("loot taken")
		return
	}
	slog.Debug("since death", "d", time.Since(e.DeathTime))
	if time.Since(e.DeathTime) < time.Millisecond*500 {
		return
	}
	e.LootDrop.Loot(p.Thing)
	e.LootDrop = nil
	e.SetAlpha(0)
}
