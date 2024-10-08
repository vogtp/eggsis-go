package thing

import (
	"fmt"
	"log/slog"
	"math"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/vogtp/eggsis-go/pkg/cfg"
	"github.com/vogtp/eggsis-go/pkg/vector"
)

type Thing struct {
	*sdl.Rect
	surface *sdl.Surface

	LP       int
	MaxLp    int
	DMG      int
	Armor    int
	Speed    int32
	MaxSpeed int32
	Gold     int

	LastAttack time.Time
	AttackFreq time.Duration

	DeathTime time.Time
}

func Create(rect sdl.Rect, imgName string) (*Thing, error) {
	t := Thing{
		Rect:       &rect,
		Speed:      cfg.BaseSpeed,
		MaxSpeed:   cfg.MaxSpeed,
		LastAttack: time.Now().Add(200 * time.Millisecond),
		AttackFreq: 100 * time.Millisecond,
	}
	if err := t.LoadImage(imgName); err != nil {
		return nil, err
	}
	return &t, nil
}

// var imgCache map[string]*sdl.Surface

func (t *Thing) LoadImage(imgName string) error {
	// if imgCache == nil {
	// 	imgCache = make(map[string]*sdl.Surface)
	// }
	// if s, ok := imgCache[imgName]; ok {
	// 	t.surface = s
	// 	return nil
	// }
	suf, err := img.Load(imgName)
	if err != nil {
		return fmt.Errorf("cannot load image %s: %w", imgName, err)
	}
	t.surface = suf
	// imgCache[imgName] = suf
	return nil
}

func (t *Thing) SetAlpha(a uint8) {
	t.surface.SetAlphaMod(a)
}

func (t Thing) String() string {
	return fmt.Sprintf("LP:%v DMG:%v Armor: %v Speed: %v", t.LP, t.DMG, t.Armor, t.Speed)
}

func (t *Thing) Move(speed vector.Speed) {
	speed.X *= t.Speed
	speed.Y *= t.Speed
	if s := speed.CalcSpeed(); s > float64(t.MaxSpeed) {
		scale := float64(t.MaxSpeed) / s
		speed.X = int32(math.Ceil(float64(speed.X) * scale))
		speed.Y = int32(math.Ceil(float64(speed.Y) * scale))
		slog.Debug("normalise speed", "speed", speed, "scale", scale)
	}
	speed.Move(t.Rect)
	t.checkBorder()
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
	// if t.IsDead() {
	// 	a := 200
	// 	if al, err := t.Surface.GetAlphaMod(); err == nil {
	// 		a = int(al) - 1
	// 	}
	// 	if a < 150 {
	// 		a = 150 // FIXME remove from array if 0
	// 	}

	// 	if err := t.Surface.SetAlphaMod(uint8(a)); err != nil {
	// 		fmt.Printf("cannot set alpha: %v", err)
	// 	}
	// }
	return t.surface.BlitScaled(nil, surf, t.Rect)
}
