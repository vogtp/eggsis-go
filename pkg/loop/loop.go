package loop

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/vogtp/eggsis-go/pkg/engine"
	vertor "github.com/vogtp/eggsis-go/pkg/vector"
)

func loop(window *sdl.Window) {
	speed = vertor.Speed{X: 0, Y: 0}
	// ttf.OpenFont()
	engine, err := engine.Create()
	if err != nil {
		panic(err)
	}
	defer engine.Free()

	windowSurface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	go events()
	color := uint32(0xff32a838)
	for running {
		if engine.Stop() {
			color = 0xffa88f32
			running = false
		}
		windowSurface.FillRect(nil, color)

		engine.Move(speed)

		if err := engine.Paint(windowSurface); err != nil {
			fmt.Printf("cannot paint: %v", err)
		}

		window.UpdateSurface()
		sdl.Delay(10)
	}
}
