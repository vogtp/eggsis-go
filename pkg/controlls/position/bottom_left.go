package position

import "github.com/veandco/go-sdl2/sdl"

func BottmLeft(x int32, y int32, w int32, h int32) *sdl.Rect {
	return &sdl.Rect{
		X: 0,
		Y: y - h,
		H: h,
		W: w,
	}
}
