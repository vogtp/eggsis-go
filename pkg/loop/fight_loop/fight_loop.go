package fight_loop

import (
	"fmt"
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
	speed         vertor.Speed
	running       = true
	stop          = false
	noSpeed       = vertor.Speed{X: 0, Y: 0}
	showVictory   = true
	processEvents = true
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
	processEvents = true
	defer func() { processEvents = false }()
	go events()
	color := uint32(0xff32a838)
	nofight := true
	var start time.Time
	fightDur := viper.GetDuration(cfg.FightDuration)
	if err := engine.StartFight(); err != nil {
		panic(err)
	}
	for running && !stop {
		loopBegin := time.Now()
		windowSurface.FillRect(nil, color)
		if err := engine.Paint(windowSurface, start); err != nil {
			slog.Error("cannot paint", "err", err)
		}
		if nofight && speed == noSpeed {
			font := fontmanager.GetFont(96)
			text, err := font.RenderUTF8Blended(fmt.Sprintf("Round %v - Move to start", engine.Round), sdl.Color{R: 0, G: 0, B: 0, A: 255})
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
		d := time.Since(loopBegin).Milliseconds()
		delay := 10
		if d < int64(delay) {
			sdl.Delay(uint32(delay) - uint32(d))
		}else{
			slog.Warn("taking TOOOO LOOOONG", "delta", d)
		}
	}
	displayVictory(window, windowSurface)
	return stop
}

func displayVictory(window *sdl.Window, windowSurface *sdl.Surface) {
	textDisp := "Victory!"
	defeat := player.Instance.LP <= 0
	if defeat {
		textDisp = "You died.... :()"
	}
	showVictory = !stop
	for !stop && showVictory {
		// slog.Warn("vicotry", "stop", stop, "showVictory", showVictory)
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
	if defeat {
		stop = true
		// os.Exit(0)
	}
}
