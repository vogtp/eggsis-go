package player

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/vogtp/eggsis-go/pkg/cfg"
	"github.com/vogtp/eggsis-go/pkg/thing"
)

type Player struct {
	*thing.Thing
}

func Create() (*Player, error) {
	p := Player{}
	r := sdl.Rect{
		X: cfg.WinX / 2,
		Y: cfg.WinY / 2,
		W: cfg.ThingSize,
		H: cfg.ThingSize * 2,
	}
	t, err := thing.Create(r, "res/drache.png")
	if err != nil {
		return nil, fmt.Errorf("cannot create base player thing: %w", err)
	}
	p.Thing = t
	p.LP = 100
	p.DMG = 5
	return &p, nil
}
