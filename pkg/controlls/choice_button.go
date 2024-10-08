package controlls

import (
	"log/slog"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/vogtp/eggsis-go/pkg/choice"
	"github.com/vogtp/eggsis-go/pkg/player"
)

type ChoiceButton struct {
	*Button
	image *sdl.Surface
}

type ChoiceList struct {
	choices []*ChoiceButton
	choice  *choice.Item
}

func (l ChoiceList) Apply(apply func(c *choice.Item)) {
	if l.choice == nil {
		slog.Error("no choice to aplly")
		return
	}
	apply(l.choice)
}

func NewChoiceList() *ChoiceList {
	return &ChoiceList{
		choices: make([]*ChoiceButton, 0),
	}
}

func NewChoiceButton(list *ChoiceList, choice *choice.Item, pos *sdl.Rect) *ChoiceButton {
	c := &ChoiceButton{}
	c.Button = NewButton(choice.Name, pos, func() {
		for _, c := range list.choices {
			c.bgColor = 233
		}

		slog.Info("chosen player", "player", player.Instance, "choice", choice.Name)
		c.bgColor = 133
		list.choice = choice
	})
	list.choices = append(list.choices, c)
	return c
}
