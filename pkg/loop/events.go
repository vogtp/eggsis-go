package loop

import (
	"fmt"
	"math"

	"github.com/veandco/go-sdl2/sdl"
	vertor "github.com/vogtp/eggsis-go/pkg/vector"
)

type vec struct {
	x, y int32
}

var (
	speed   vertor.Speed
	running = true
)

func events() {
	fmt.Println("Starting event loop")
	defer fmt.Println("Stopping event loop")
	for running {
		event := sdl.PollEvent()
		if event == nil {
			continue
		}
		switch e := event.(type) {
		case *sdl.QuitEvent:
			//fmt.Printf("Got quit event %v\n", e)
			running = false
			break
		case *sdl.KeyboardEvent:

			switch e.Keysym.Sym {
			case sdl.K_LEFT:
				speed.X = -1 * calcSpeed(e, speed.X)
				//speed.y = 0
			case sdl.K_RIGHT:
				speed.X = calcSpeed(e, speed.X)
				//speed.y = 0
			case sdl.K_UP:
				//speed.x = 0
				speed.Y = -1 * calcSpeed(e, speed.Y)
			case sdl.K_DOWN:
				//speed.x = 0
				speed.Y = calcSpeed(e, speed.Y)
			case sdl.K_END:
				running = false
			default:
				//fmt.Printf("%#v\n", e)
			}
		case *sdl.MouseMotionEvent:
			// speed.x = e.X
			// speed.y = e.Y
		default:
			//fmt.Printf("%#v\n", e)
		}
	}
}

func calcSpeed(e *sdl.KeyboardEvent, s int32) int32 {
	switch e.State {
	case sdl.RELEASED:
		return 0
	case sdl.PRESSED:
		s = int32(math.Abs(float64(s)))
		// s++
		// if s > cfg.MaxSpeed {
		// 	s = cfg.MaxSpeed
		// }
		return s
	default:
		return 0
	}
}
