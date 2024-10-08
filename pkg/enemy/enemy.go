package enemy

import (
	"fmt"
	"math/rand/v2"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/vogtp/eggsis-go/pkg/cfg"
	"github.com/vogtp/eggsis-go/pkg/loot"
	"github.com/vogtp/eggsis-go/pkg/thing"
)

type Enemy struct {
	*thing.Thing
	LootDrop loot.Loot
}

func Create(t *thing.Thing, round int) (*Enemy, error) {

	r := randRect()
	for t.HasIntersection(&r) {
		r = randRect()
	}
	t, err := thing.Create(r, "res/meat.png")
	if err != nil {
		return nil, fmt.Errorf("cannot create base enemy thing: %w", err)
	}
	e := Enemy{Thing: t}
	e.DMG = rand.IntN(1) + 4*round
	e.LP = rand.IntN(50) + 50*round
	e.Speed = rand.Int32N(e.Speed) + 1*int32(round)
	if rand.IntN(2) == 1 {
		e.LootDrop = loot.Gold(rand.IntN(e.DMG*5) + e.DMG)
	} else {
		e.LootDrop = loot.Heal(rand.IntN(e.DMG*5) + e.DMG)
	}
	return &e, nil
}

func randRect() sdl.Rect {
	x := rand.Int32N(cfg.WinX - cfg.ThingSize)
	y := rand.Int32N(cfg.WinY - cfg.ThingSize)
	r := sdl.Rect{
		X: x,
		Y: y,
		W: cfg.ThingSize,
		H: cfg.ThingSize + 15,
	}
	return r
}
