package controlls

import (
	"log/slog"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/vogtp/eggsis-go/pkg/choice"
	"github.com/vogtp/eggsis-go/pkg/player"
)

type ChoiceButton struct {
	*Button
	image  *sdl.Surface
	choice *choice.Item
}

type ChoiceList struct {
	choices         []*ChoiceButton
	choiceSelection map[*ChoiceButton]bool
}

func (l ChoiceList) Apply(apply func(c *choice.Item)) {
	if l.choiceSelection == nil {
		slog.Error("no choice to aplly")
		return
	}
	for c, ok := range l.choiceSelection {
		if ok {
			slog.Debug("Appling choice", "choice", c.choice)
			apply(c.choice)
		}
	}
}

func NewChoiceList() *ChoiceList {
	return &ChoiceList{
		choices:         make([]*ChoiceButton, 0),
		choiceSelection: make(map[*ChoiceButton]bool),
	}
}

func NewChoiceButton(list *ChoiceList, choice *choice.Item, pos *sdl.Rect, multi bool) *ChoiceButton {
	c := &ChoiceButton{
		choice: choice,
	}
	c.Button = NewButton(choice.Name, pos, func() {
		if !multi {
			for k := range list.choiceSelection {
				list.choiceSelection[k] = false
			}
		}
		list.choiceSelection[c] = !list.choiceSelection[c]
		for _, pos := range list.choices {
			pos.bgColor = 233
			if list.choiceSelection[pos] {
				pos.bgColor = 133
			}
		}

		slog.Info("chosen mod", "player", player.Instance, "choice", choice.Name)
		
	})
	list.choices = append(list.choices, c)
	list.choiceSelection[c] = false

	return c
}
