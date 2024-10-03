package thing

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/vogtp/eggsis-go/pkg/cfg"
	vertor "github.com/vogtp/eggsis-go/pkg/vector"
)

type Thing struct {
	*sdl.Rect
	surface *sdl.Surface

	LP  int
	DMG int
	Armor int
}

func Create(rect sdl.Rect, imgName string) (*Thing, error) {
	t := Thing{
		Rect: &rect,
	}
	suf, err := img.Load(imgName)
	if err != nil {
		return nil, fmt.Errorf("cannot load image %s: %w", imgName, err)
	}
	t.surface = suf
	return &t, nil
}

func (t *Thing) Move(speed vertor.Speed) {
	speed.Move(t.Rect)
	t.checkBorder()
}

func (t *Thing) IsDead() bool {
	d := t.LP < 0
	if d && t.LP != -255{
		t.LP = -255
		if suf, err := img.Load("res/ball.png"); err == nil {
			suf.SetAlphaMod(200)
			t.surface = suf
		} else {
			fmt.Printf("cannot load dead img: %v\n", err)
		}
	}
	return d
}

func (t *Thing) checkBorder() {
	if t.Rect.X < 0 {
		t.Rect.X = 0
	}
	if t.Rect.Y < 0 {
		t.Rect.Y = 0
	}
	if t.Rect.X+t.Rect.W > cfg.WinX {
		t.Rect.X = cfg.WinX - t.Rect.W
	}
	if t.Rect.Y+t.Rect.H > cfg.WinY {
		t.Rect.Y = cfg.WinY - t.Rect.H
	}
}

func (t *Thing) Free() {
	if t.surface != nil {
		t.surface.Free()
	}
}

func (t *Thing) Paint(surf *sdl.Surface) error {
	if t.IsDead() {
		a := 200
		if al, err := t.surface.GetAlphaMod(); err == nil{
			a = int(al) - 2
		}
		if a < 0 {
			a = 0 // FIXME remove from array if 0
		}

		if err:=t.surface.SetAlphaMod(uint8(a)); err!=nil{
			fmt.Printf("cannot set alpha: %v",err)
		}
	}
	return t.surface.BlitScaled(nil, surf, t.Rect)
}
