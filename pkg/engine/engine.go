package engine

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/spf13/viper"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"github.com/vogtp/eggsis-go/pkg/cfg"
	"github.com/vogtp/eggsis-go/pkg/choice"
	"github.com/vogtp/eggsis-go/pkg/enemy"
	"github.com/vogtp/eggsis-go/pkg/fontmanager"
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
	Round         int
}

var instance *Engine

func Create() (*Engine, error) {
	e := instance
	if instance == nil {
		e = &Engine{}
		instance = e
	}

	if e.player == nil {
		p, err := player.Create()
		if err != nil {
			return nil, fmt.Errorf("cannot create player: %w", err)
		}
		e.player = p
	}

	e.font = fontmanager.GetFont(18)

	return e, nil
}

func (e *Engine) StartFight() error {
	e.Round++
	slog.Warn("Starting fight", "round", e.Round)
	e.player.SetToStart()
	e.enemies = make([]*enemy.Enemy, EnemyCnt)
	for i := 0; i < EnemyCnt; i++ {
		en, err := enemy.Create(e.player.Thing, e.Round)
		if err != nil {
			return err
		}
		e.enemies[i] = en
	}
	return nil
}

func (e *Engine) CreatePlayer(c *choice.Item) error {
	if player.Instance != nil {
		return nil
	}
	p, err := player.Create()
	if err != nil {
		return fmt.Errorf("cannot create player: %w", err)
	}
	e.player = p

	e.player.Name = c.Name
	if err := e.player.LoadImage(c.Image); err != nil {
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

func (e *Engine) Paint(surf *sdl.Surface, start time.Time) error {
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
	if !start.IsZero() {
		s["Duration"] = fmt.Sprintf("%s", viper.GetDuration(cfg.FightDuration).Truncate(time.Second)-time.Since(start).Truncate(time.Second))
	}
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
