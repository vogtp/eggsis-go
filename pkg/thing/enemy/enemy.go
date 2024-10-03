package enemy

import (
	"fmt"
	"math"
	"math/rand/v2"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/vogtp/eggsis-go/pkg/cfg"
	"github.com/vogtp/eggsis-go/pkg/thing"
	"github.com/vogtp/eggsis-go/pkg/thing/player"
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

func calcDmg(dmg int, armor int) int {
	d := dmg - armor
	if d < 0 {
		d = 0
	}
	return d
}

func (e *Enemy) Fight(p *player.Player) {
	p.LP -= calcDmg(e.DMG, p.Armor)
	e.LP -= calcDmg(p.DMG , e.Armor)
}

func (e *Enemy) MoveTo(p *player.Player, others []Enemy) {
	if p.HasIntersection(e.Rect) {
		e.Fight(p)
		return
	}
	x := p.X - e.Rect.X
	y := p.Y - e.Rect.Y
	l := math.Sqrt(float64(x*x + y*y))
	dx := int32(math.Round(float64(x)/l)) * e.speed
	dy := int32(math.Round(float64(y)/l)) * e.speed
	e.X += dx
	e.Y += dy
	for _, o := range others {
		if o == *e {
			continue
		}
		if o.HasIntersection(e.Rect) && !o.IsDead() {
			e.X -= dx
			e.Y -= dy
			return
		}
	}
}

func randRect() sdl.Rect {
	x := rand.Int32N(cfg.WinX - cfg.ThingSize)
	y := rand.Int32N(cfg.WinY - cfg.ThingSize)
	r := sdl.Rect{
		X: x,
		Y: y,
		W: cfg.ThingSize - 10,
		H: cfg.ThingSize,
	}
	return r
}
