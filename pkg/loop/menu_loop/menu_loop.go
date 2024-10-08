package menuloop

import (
	"log/slog"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/vogtp/eggsis-go/pkg/controlls"
	"github.com/vogtp/eggsis-go/pkg/controlls/position"
	"github.com/vogtp/eggsis-go/pkg/player"
	vertor "github.com/vogtp/eggsis-go/pkg/vector"
)

var (
	speed   vertor.Speed
	running = true
	stop    = false
	buttons []controlls.Clickable
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
	defer func() { buttons = buttons[:0] }()
	buttons = append(buttons, controlls.NewButton("Quit", position.BottmLeft(windowSurface.W, windowSurface.H, 100, 30), func() { stop = true }))
	buttons = append(buttons, controlls.NewButton("Start Fight", position.BottmRight(windowSurface.W, windowSurface.H, 100, 30), func() { running = false }))
	
	choose:=modsMenu()
	defer choose()
	if player.Instance == nil {
		slog.Info("Player Menu", "player", player.Instance)
		choose := playerMenu()
		defer choose()
	}
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
