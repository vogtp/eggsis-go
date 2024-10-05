package loop

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/vogtp/eggsis-go/pkg/cfg"
	"github.com/vogtp/eggsis-go/pkg/loop/fight_loop"
	menuloop "github.com/vogtp/eggsis-go/pkg/loop/menu_loop"
)

func Start() {
	sdl.Main(run)
}

func run() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow(cfg.AppName, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		cfg.WinX, cfg.WinY, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	sdl.Do(func() { loop(window) })

}

type LoopFunc func(*sdl.Window) bool

var loopFuncs []LoopFunc = []LoopFunc{menuloop.Run, fight_loop.Run}

func loop(window *sdl.Window) {
	for {
		for _, loopFunc := range loopFuncs {
			if loopFunc(window) {
				return
			}
			sdl.Delay(100)
		}
	}
}
