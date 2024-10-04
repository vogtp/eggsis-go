package vector

import "github.com/veandco/go-sdl2/sdl"

type Speed struct {
	X int32
	Y int32
}

func (s Speed) Move(r *sdl.Rect) {
	r.X += s.X
	r.Y += s.Y
}
