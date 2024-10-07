package controlls

import "github.com/veandco/go-sdl2/sdl"

type Clickable interface {
	IsClicked(*sdl.Rect) bool
	Paint(*sdl.Surface) error
	Action()
}
