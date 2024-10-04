package enemy

import (
	"fmt"
	"math/rand/v2"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/vogtp/eggsis-go/pkg/cfg"
	"github.com/vogtp/eggsis-go/pkg/thing"
)

type Enemy struct {
	*thing.Thing
	speed int32
}

func Create(t *thing.Thing) (*Enemy, error) {
	e := Enemy{
		speed: cfg.BaseSpeed,
	}
	r := randRect()
	for t.HasIntersection(&r) {
		r = randRect()
	}
	t, err := thing.Create(r, "res/meat.png")
	if err != nil {
		return nil, fmt.Errorf("cannot create base enemy thing: %w", err)
	}
	e.Thing = t
	e.DMG = rand.IntN(2) + 1
	e.LP = rand.IntN(50) + 50
	return &e, nil
}

func randRect() sdl.Rect {
	x := rand.Int32N(cfg.WinX - cfg.ThingSize)
	y := rand.Int32N(cfg.WinY - cfg.ThingSize)
	r := sdl.Rect{
		X: x,
		Y: y,
		W: cfg.ThingSize,
		H: cfg.ThingSize+15,
	}
	return r
}
