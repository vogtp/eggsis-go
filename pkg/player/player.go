package player

import (
	"fmt"
	"log/slog"

	"github.com/spf13/viper"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/vogtp/eggsis-go/pkg/cfg"
	"github.com/vogtp/eggsis-go/pkg/thing"
)

type Egg struct {
	*thing.Thing
	Name  string
	MaxLp int
	Gold  int
}

//FIXME stupid hack. but player creation depends on instance beeing visible...
var Instance *Egg

func Create() (*Egg, error) {
	slog.Info("player create","instance", Instance)
	if Instance != nil {
		Instance.SetToStart()
		return Instance, nil
	}
	p := Egg{
		Gold: 50,
	}
	Instance = &p
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
	p.Armor = 2
	p.MaxLp = p.LP
	return &p, nil
}

func (p Egg) String() string{
	return fmt.Sprintf("%s: %s", p.Name, p.Thing.String())
}

func (p *Egg) SetToStart() {
	p.LP = Instance.MaxLp
	p.X = cfg.WinX / 2
	p.Y = cfg.WinY / 2

}

func (p *Egg) IsDead() bool {
	d := p.Thing.IsDead()

	if d && !viper.GetBool(cfg.PlayerDeath) {
		slog.Debug("No player death")
		t, _ := thing.Create(*p.Rect, "res/egg.png")
		p.Thing = t
		return false
	}
	return d
}
