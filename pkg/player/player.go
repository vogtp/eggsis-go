package player

import (
	"fmt"
	"log/slog"

	"github.com/spf13/viper"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/vogtp/eggsis-go/pkg/cfg"
	"github.com/vogtp/eggsis-go/pkg/thing"
)

type Player struct {
	*thing.Thing
	MaxLp int
	Gold  int
}

var player *Player

func Create() (*Player, error) {
	if player != nil {
		return player, nil
	}
	p := Player{}
	player = p
	r := sdl.Rect{
		X: cfg.WinX / 2,
		Y: cfg.WinY / 2,
		W: cfg.ThingSize,
		H: cfg.ThingSize + 15,
	}
	t, err := thing.Create(r, "res/egg.png")
	if err != nil {
		return nil, fmt.Errorf("cannot create base player thing: %w", err)
	}
	p.Thing = t
	p.LP = viper.GetInt(cfg.PlayerLP)
	p.DMG = 3
	p.Armor = 1
	p.MaxLp = p.LP
	return &p, nil
}

func (p *Player) IsDead() bool {
	d := p.Thing.IsDead()

	if d && !viper.GetBool(cfg.PlayerDeath) {
		slog.Debug("No player death")
		t, _ := thing.Create(*p.Rect, "res/egg.png")
		p.Thing = t
		return false
	}
	return d
}
