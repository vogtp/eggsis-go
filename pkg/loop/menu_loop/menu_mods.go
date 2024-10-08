package menuloop

import (
	"log/slog"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/vogtp/eggsis-go/pkg/choice"
	"github.com/vogtp/eggsis-go/pkg/controlls"
	"github.com/vogtp/eggsis-go/pkg/player"
)

func modsMenu() func() {
	list := controlls.NewChoiceList()
	cb := make([]*controlls.ChoiceButton, 0)
	for i, c := range choice.Mods {
		w := int32(200)
		r := sdl.Rect{
			X: 10*int32(i) + int32(i)*w,
			Y: 220,
			W: w,
			H: 100,
		}
		cb = append(cb, controlls.NewChoiceButton(list, &c, &r))
	}
	for _, c := range cb {
		buttons = append(buttons, c)
	}
	choose := func() {
		list.Apply(func(c *choice.Item) {
			slog.Info("Appling mod", "player", player.Instance, "mod", c)
			c.Modifier(player.Instance)
		})
	}
	return choose
}
