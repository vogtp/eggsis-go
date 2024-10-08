package engine

import (
	"log/slog"
	"time"

	"github.com/vogtp/eggsis-go/pkg/enemy"
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

func (e *Engine) Fight(player *player.Egg, enemy *enemy.Enemy) {
	slog.Debug("fight", "enemy", enemy, "player", player)
	if enemy.IsDead() {
		player.GetLooted(enemy)
		return
	}
	player.LP -= calcDmg(enemy.Thing, player.Thing)
	enemy.LP -= calcDmg(player.Thing, enemy.Thing)
	if enemy.IsDead() {
		enemy.DeathTime = time.Now()
	}
}
