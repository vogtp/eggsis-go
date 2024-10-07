package engine

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"github.com/vogtp/eggsis-go/pkg/enemy"
	"github.com/vogtp/eggsis-go/pkg/fontmanager"
	"github.com/vogtp/eggsis-go/pkg/fontmanager/choice"
	"github.com/vogtp/eggsis-go/pkg/player"
)

const (
	EnemyCnt   = 10
	EnemySpwan = 100
)

type Engine struct {
	player        *player.Egg
	enemies       []*enemy.Enemy
	enemySpawnCnt int
	font          *ttf.Font
}

func Create() (*Engine, error) {
	e := Engine{}

	p, err := player.Create()
	if err != nil {
		return nil, fmt.Errorf("cannot create player: %w", err)
	}
	e.player = p
	e.enemies = make([]*enemy.Enemy, EnemyCnt)
	for i := 0; i < EnemyCnt; i++ {
		en, err := enemy.Create(e.player.Thing)
		if err != nil {
			return nil, err
		}
		e.enemies[i] = en
	}
	e.font = fontmanager.GetFont(18)

	return &e, nil
}

func (e *Engine) CreatePlayer(c choice.Item) error {
	if len(e.player.Name)>0 {
		return nil
	}
	e.player.Name =  c.Name
	if err:=e.player.LoadImage(c.Image); err != nil{
		return err
	}
	c.Modifier(e.player)
	return nil
}

func (e *Engine) Stop() bool {
	if e.player.IsDead() {
		return true
	}

	return false
}

func (e *Engine) Paint(surf *sdl.Surface) error {
	// e.removeEnemies()
	e.paintPlayerHealth(surf)
	if err := e.player.Paint(surf); err != nil {
		return err
	}
	for _, en := range e.enemies {
		if err := en.Paint(surf); err != nil {
			return err
		}
	}
	s := e.player.Stats()
	s["Enemies"] = fmt.Sprintf("%v", len(e.enemies))
	e.paintStats(surf, s)
	return nil
}

func (e *Engine) Free() {
	// if e.player != nil {
	// 	e.player.Free()
	// }
	if e.font != nil {
		e.font.Close()
	}
}
