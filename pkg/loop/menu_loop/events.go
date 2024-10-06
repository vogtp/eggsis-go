package menuloop

import (
	"log/slog"
	"math"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/vogtp/eggsis-go/pkg/cfg"
)

func events() {
	slog.Warn("Starting event loop")
	defer slog.Warn("Stopping event loop")
	for running {
		event := sdl.PollEvent()
		if event == nil {
			continue
		}
		switch e := event.(type) {
		case *sdl.QuitEvent:
			//fmt.Printf("Got quit event %v\n", e)
			stop = true
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
				if e.State != sdl.RELEASED {
					continue
				}
				running = false
			case sdl.K_q:
				stop = true
			default:
				//fmt.Printf("%#v\n", e)
			}
		case *sdl.MouseButtonEvent:
			r := sdl.Rect{X: e.X, Y: e.Y, H: 1, W: 1}
			for _, b := range buttons {
				if b.IsClicked(&r) {
					break
				}
			}
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
		s++
		if s > cfg.MaxSpeed {
			s = cfg.MaxSpeed
		}
		return s
	default:
		return 0
	}
}
