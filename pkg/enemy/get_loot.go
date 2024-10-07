package enemy

import (
	"log/slog"
	"time"

	"github.com/vogtp/eggsis-go/pkg/player"
)

func (e *Enemy) GetLooted(p *player.Egg) {
	if e.LootDrop == nil {
		slog.Debug("loot taken")
		return
	}
	slog.Debug("since death", "d", time.Since(e.DeathTime))
	if time.Since(e.DeathTime) < time.Millisecond*500 {
		return
	}
	e.LootDrop.Loot(p)
	e.LootDrop = nil
	e.Surface.SetAlphaMod(0)
}
