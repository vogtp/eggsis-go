package controlls

import (
	"log/slog"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/vogtp/eggsis-go/pkg/cfg"
	"github.com/vogtp/eggsis-go/pkg/fontmanager"
)

type ActionFunc func()

type Button struct {
	buttonRect   *sdl.Rect
	labelRect *sdl.Rect

	label   string
	textSurf *sdl.Surface

	action ActionFunc
}

func NewButton(label string, action ActionFunc) *Button {
	b := Button{
		label:  label,
		action: action,
	}
	font := fontmanager.GetFont(18)
	text, err := font.RenderUTF8Blended(label, sdl.Color{R: 255, G: 0, B: 0, A: 255})
	if err != nil {
		panic(err)
	}
	rText := &sdl.Rect{X: cfg.WinX/2 - text.W/2, Y: cfg.WinY/2 - text.H/2, W: text.W, H: text.H}
	b.buttonRect = &sdl.Rect{X: rText.X - 5, Y: rText.Y - 5, W: text.W + 8, H: text.H + 8}
	b.labelRect = rText
	b.textSurf = text
	return &b
}

func (b Button) Paint(surf *sdl.Surface) error {
	if err := surf.FillRect(b.buttonRect, 233); err != nil {
		return err
	}
	if err := b.textSurf.Blit(nil, surf, b.labelRect); err != nil {
		return err
	}
	return nil
}

func (b *Button) IsClicked(rct *sdl.Rect) bool {
	if rct.H != 1 {
		rct.H = 1
	}
	if rct.W != 1 {
		rct.W = 1
	}
	slog := slog.Default().With("label", b.label)
	slog.Debug("click", "loc", rct, "bu", b.buttonRect)
	clicked := b.buttonRect.HasIntersection(rct)
	if clicked {
		slog.Debug("button click")
		b.action()
	}
	return clicked
}
