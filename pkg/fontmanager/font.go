package fontmanager

import (
	"fmt"

	"github.com/veandco/go-sdl2/ttf"
)

var fontInit bool

func GetFont(size int) *ttf.Font {
	if !fontInit {

		if err := ttf.Init(); err != nil {
			panic(fmt.Errorf("cannot initialise TTF subsystem: %w", err))
		}
	}
	f, err := ttf.OpenFont("res/Go-Regular.ttf", size)
	if err != nil {
		panic(fmt.Errorf("font error: %w", err))
	}
	return f
}
