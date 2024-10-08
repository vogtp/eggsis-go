package menuloop

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/vogtp/eggsis-go/pkg/choice"
	"github.com/vogtp/eggsis-go/pkg/controlls"
	"github.com/vogtp/eggsis-go/pkg/engine"
	"github.com/vogtp/eggsis-go/pkg/player"
)

func playerMenu() func() {
	engine, err := engine.Create()
	if err != nil {
		panic(err)
	}
	defer engine.Free()
	list := controlls.NewChoiceList()
	cb := make([]*controlls.ChoiceButton, 0)
	for i, c := range choice.Players {
		w := int32(200)
		r := sdl.Rect{
			X: 10*int32(i) + int32(i)*w,
			Y: 100,
			W: w,
			H: 100,
		}
		cb = append(cb, controlls.NewChoiceButton(list, &c, &r, false))
	}
	for _, c := range cb {
		buttons = append(buttons, c)
	}
	choose := func() {
		list.Apply(func(c *choice.Item) {
			player.Instance = nil // FIXME really stupid hack
			engine.CreatePlayer(c)
		})
	}
	return choose
}
