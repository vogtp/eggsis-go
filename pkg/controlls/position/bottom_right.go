package position

import "github.com/veandco/go-sdl2/sdl"

func BottmRight(x int32, y int32, w int32, h int32) *sdl.Rect {
	return &sdl.Rect{
		X: x -w,
		Y: y - h,
		H: h,
		W:w,
	}
}