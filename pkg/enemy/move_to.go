package enemy

import (
	"math"

	"github.com/vogtp/eggsis-go/pkg/player"
)

func (e *Enemy) MoveTo(p *player.Player, others []Enemy) {
	if p.HasIntersection(e.Rect) {
		e.Fight(p)
		return
	}
	x := p.X - e.Rect.X
	y := p.Y - e.Rect.Y
	l := math.Sqrt(float64(x*x + y*y))
	dx := int32(math.Round(float64(x)/l)) * e.Speed
	dy := int32(math.Round(float64(y)/l)) * e.Speed
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
