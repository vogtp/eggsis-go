package menuloop

import (
	"github.com/veandco/go-sdl2/sdl"
	vertor "github.com/vogtp/eggsis-go/pkg/vector"
)

var (
	speed   vertor.Speed
	running = true
)

func Run(window *sdl.Window) bool {
	stop := false
	running = true
	speed = vertor.Speed{X: 0, Y: 0}

	windowSurface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	go events()
	//	color := uint32(0xff32a838)
	for running {
		windowSurface.FillRect(nil, 0)

		window.UpdateSurface()
		sdl.Delay(10)
	}
	return stop
}
