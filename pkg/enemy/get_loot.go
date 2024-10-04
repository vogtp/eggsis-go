package enemy

import (
	"log/slog"
	"time"

	"github.com/vogtp/eggsis-go/pkg/player"
)

func (e *Enemy) GetLooted(p *player.Player) {
	if e.lootDrop == nil {
		slog.Debug("loot taken")
		return
	}
	slog.Debug("since death", "d", time.Since(e.DeathTime))
	if time.Since(e.DeathTime) < time.Millisecond*500 {
		return
	}
	l := *e.lootDrop
	l(p)
	e.lootDrop = nil
	e.Surface.SetAlphaMod(0)
}
