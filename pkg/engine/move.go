package engine

import (
	"github.com/vogtp/eggsis-go/pkg/enemy"
	"github.com/vogtp/eggsis-go/pkg/vector"
)

func (e *Engine) Move(s vector.Speed) {
	e.player.Move(s)
	e.enemySpawnCnt++
	if e.enemySpawnCnt > EnemySpwan {
		e.enemySpawnCnt = 0
		if en, err := enemy.Create(e.player.Thing, e.Round); err == nil {
			e.enemies = append(e.enemies, en)
		}
	}
	for _, en := range e.enemies {
		if en.IsDead() {
			// fmt.Println("enemy dead")
			//slices.Delete(e.enemies, i, i+1)
			//continue
		}
		en.MoveTo(e.player, e.enemies)
	}
	e.player.Action(e.enemies)
}
