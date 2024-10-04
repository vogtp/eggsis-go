package thing

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
)

func (t *Thing) IsDead() bool {
	d := t.LP < 0
	if d && t.LP != -255 {
		t.LP = -255
		if suf, err := img.Load("res/meat_dead.png"); err == nil {
			suf.SetAlphaMod(200)
			t.surface = suf
		} else {
			fmt.Printf("cannot load dead img: %v\n", err)
		}
	}
	return d
}
