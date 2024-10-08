package fight_loop

import (
	"log/slog"
	"time"

	"github.com/spf13/viper"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/vogtp/eggsis-go/pkg/cfg"
	"github.com/vogtp/eggsis-go/pkg/engine"
	"github.com/vogtp/eggsis-go/pkg/fontmanager"
	"github.com/vogtp/eggsis-go/pkg/player"
	vertor "github.com/vogtp/eggsis-go/pkg/vector"
)

var (
	speed   vertor.Speed
	running = true
	stop    = false
	noSpeed = vertor.Speed{X: 0, Y: 0}
	showVictory = true
)

func Run(window *sdl.Window) bool {
	running = true
	stop = false
	speed = noSpeed

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
	nofight := true
	var start time.Time
	fightDur := viper.GetDuration(cfg.FightDuration)
	for running && !stop {
		windowSurface.FillRect(nil, color)
		if err := engine.Paint(windowSurface, start); err != nil {
			slog.Error("cannot paint", "err", err)
		}
		if nofight && speed == noSpeed {
			font := fontmanager.GetFont(96)
			text, err := font.RenderUTF8Blended("Move to start", sdl.Color{R: 0, G: 0, B: 0, A: 255})
			if err != nil {
				panic(err)
			}

			// Draw the text around the center of the window
			if err = text.Blit(nil, windowSurface, &sdl.Rect{X: cfg.WinX/2 - text.W/2, Y: cfg.WinY/2 - text.H/2, W: 0, H: 0}); err != nil {
				panic(err)
			}
			window.UpdateSurface()
			sdl.Delay(10)
			continue
		}

		if start.IsZero() {
			start = time.Now()
		}
		nofight = false

		if time.Since(start) > fightDur {
			break
		}
		if engine.Stop() {
			color = 0xffa88f32
			running = false
		}
		engine.Move(speed)

		window.UpdateSurface()
		sdl.Delay(10)
	}
	displayVictory(window, windowSurface)
	return stop
}

func displayVictory(window *sdl.Window, windowSurface *sdl.Surface) {
	textDisp := "Victory!"
	if player.Instance.LP <= 0 {
		textDisp = "You died.... :()"
	}
	running = true
	showVictory = true
	//	speed = noSpeed
	for !stop && showVictory {
		font := fontmanager.GetFont(96)
		text, err := font.RenderUTF8Blended(textDisp, sdl.Color{R: 0, G: 0, B: 0, A: 255})
		if err != nil {
			panic(err)
		}

		// Draw the text around the center of the window
		if err = text.Blit(nil, windowSurface, &sdl.Rect{X: cfg.WinX/2 - text.W/2, Y: cfg.WinY/2 - text.H/2, W: 0, H: 0}); err != nil {
			panic(err)
		}
		window.UpdateSurface()
		sdl.Delay(10)
	}

}
