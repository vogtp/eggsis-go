package menuloop

import (
	"log/slog"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/vogtp/eggsis-go/pkg/controlls"
	"github.com/vogtp/eggsis-go/pkg/controlls/position"
	"github.com/vogtp/eggsis-go/pkg/engine"
	"github.com/vogtp/eggsis-go/pkg/fontmanager/choice"
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
	defer func() {buttons = buttons[:0]}()
	buttons = append(buttons, controlls.NewButton("Quit", position.BottmLeft(windowSurface.W, windowSurface.H, 100, 30), func() { stop = true }))
	buttons = append(buttons, controlls.NewButton("Start Fight", position.BottmRight(windowSurface.W, windowSurface.H, 100, 30), func() { running = false }))
	if player.Instance == nil {
		slog.Info("Player Menu", "player", player.Instance)
		playerMenu()
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

func playerMenu() {
	engine, err := engine.Create()
	if err != nil {
		panic(err)
	}
	defer engine.Free()

	cb := make([]*controlls.ChoiceButton, 0)
	for i, c := range choice.Players {
		w := int32(200)
		r := sdl.Rect{
			X: 10*int32(i) + int32(i)*w,
			Y: 100,
			W: w,
			H: 100,
		}
		cb = append(cb, controlls.NewChoiceButton(c, &r, engine))
	}
	for _, c := range cb {
		buttons = append(buttons, c)
	}
}
