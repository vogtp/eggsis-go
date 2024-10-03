package engine

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"github.com/vogtp/eggsis-go/pkg/cfg"
	"github.com/vogtp/eggsis-go/pkg/thing/enemy"
	"github.com/vogtp/eggsis-go/pkg/thing/player"
	vertor "github.com/vogtp/eggsis-go/pkg/vector"
)

const (
	EnemyCnt   = 10
	EnemySpwan = 100
)

type Engine struct {
	player        *player.Player
	enemies       []enemy.Enemy
	enemySpawnCnt int
	font *ttf.Font
}

func Create() (*Engine, error) {
	e := Engine{}

	p, err := player.Create()
	if err != nil {
		return nil, fmt.Errorf("cannot create player: %w", err)
	}
	e.player = p
	e.enemies = make([]enemy.Enemy, EnemyCnt)
	for i := 0; i < EnemyCnt; i++ {
		en, err := enemy.Create(e.player.Thing)
		if err != nil {
			return nil, err
		}
		e.enemies[i] = *en
	}
	// if font, err = ttf.OpenFont(fontPath, fontSize); err != nil {
	// 	return err
	// }

	// defer font.Close()
	return &e, nil
}

func (e *Engine) Stop() bool {
	if e.player.IsDead() {
		return true
	}

	return false
}

func (e *Engine) Move(s vertor.Speed) {
	e.player.Move(s)
	e.enemySpawnCnt++
	if e.enemySpawnCnt > EnemySpwan {
		e.enemySpawnCnt = 0
		if en, err := enemy.Create(e.player.Thing); err == nil {
			e.enemies = append(e.enemies, *en)
		}
	}
	for _, en := range e.enemies {
		if en.IsDead() {
			fmt.Println("enemy dead")
			//slices.Delete(e.enemies, i, i+1)
			continue
		}
		en.MoveTo(e.player, e.enemies)
	}
}

func (e Engine) paintPlayerHealth(surf *sdl.Surface) {
	y := cfg.WinX / e.player.MaxLp * e.player.LP
	surf.FillRect(&sdl.Rect{X: 1, Y: 1, H: 20, W: cfg.WinX}, 0x03fcdb)
	surf.FillRect(&sdl.Rect{X: 1, Y: 1, H: 20, W: int32(y)}, 0xfc2403)
}

func (e *Engine) Paint(surf *sdl.Surface) error {
	e.paintPlayerHealth(surf)
	if err := e.player.Paint(surf); err != nil {
		return err
	}
	for _, en := range e.enemies {
		if err := en.Paint(surf); err != nil {
			return err
		}
	}
	return nil
}

func (e *Engine) Free() {
	if e.player != nil {
		e.player.Free()
	}
}
