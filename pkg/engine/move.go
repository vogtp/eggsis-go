package engine

import (
	"github.com/vogtp/eggsis-go/pkg/enemy"
	"github.com/vogtp/eggsis-go/pkg/vector"
)

func (e *Engine) cleanupEnemies() {
	if e.deadEnemies < 10 {
		return
	}
	e2 := e.enemies
	e.enemies = make([]*enemy.Enemy, 0, len(e2))
	for _, en := range e2 {
		if en.Delete {
			en.Free()
		} else {
			e.enemies = append(e.enemies, en)
		}
	}
	e.deadEnemies = 0
}

func (e *Engine) Move(s vector.Speed) {
	e.cleanupEnemies()
	e.player.Move(s)
	go e.player.Action(e.enemies)
	e.enemySpawnCnt++
	if e.enemySpawnCnt > EnemySpwan {
		e.enemySpawnCnt = 0
		if en, err := enemy.Create(e.player.Thing, e.Round); err == nil {
			e.enemies = append(e.enemies, en)
		}
	}
	for _, en := range e.enemies {
		// if en.IsDead() {
		// 	// fmt.Println("enemy dead")
		// 	//slices.Delete(e.enemies, i, i+1)
		// 	//continue
		// }
		if e.player.HasIntersection(en.Rect) {
			e.Fight(e.player, en)
			continue
		}
		en.MoveTo(e.player.Rect, e.enemies)
	}
}
