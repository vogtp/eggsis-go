package thing

import (
	"log/slog"

	"github.com/veandco/go-sdl2/img"
)

func (t *Thing) IsDead() bool {
	d := t.LP <= 0

	if d && t.LP != -255 {
		t.LP = -255
		if suf, err := img.Load("res/gold.png"); err == nil {
			suf.SetAlphaMod(250)
			t.Surface = suf
		} else {
			slog.Warn("cannot load dead img", "err", err)
		}
	}
	return d
}
