package menuloop

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/vogtp/eggsis-go/pkg/controlls"
	"github.com/vogtp/eggsis-go/pkg/controlls/position"
	vertor "github.com/vogtp/eggsis-go/pkg/vector"
)

var (
	speed   vertor.Speed
	running = true
	stop    = false
	buttons []*controlls.Button
)

func Run(window *sdl.Window) bool {
	stop = false
	running = true
	speed = vertor.Speed{X: 0, Y: 0}

	windowSurface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	go events()
	//	color := uint32(0xff32a838)
	buttons = append(buttons, controlls.NewButton("Quit", position.BottmLeft(windowSurface.W, windowSurface.H, 100, 30), func() { stop = true }))
	buttons = append(buttons, controlls.NewButton("Start Fight", position.BottmRight(windowSurface.W, windowSurface.H, 100, 30), func() { running = false }))
	for running && !stop {
		windowSurface.FillRect(nil, 0)

		for _, b := range buttons {
			b.Paint(windowSurface)
		}

		window.UpdateSurface()
		sdl.Delay(10)
	}
	return stop
}
