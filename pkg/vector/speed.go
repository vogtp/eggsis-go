package vector

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type Speed struct {
	X int32
	Y int32
}

func (s Speed) Move(r *sdl.Rect) {
	r.X += s.X
	r.Y += s.Y
}

func (s Speed) CalcSpeed() float64{
	return math.Sqrt(float64(s.X*s.X + s.Y*s.Y))
}