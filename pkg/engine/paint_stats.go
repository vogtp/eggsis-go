package engine

import (
	"fmt"
	"sort"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/vogtp/eggsis-go/pkg/cfg"
)

func (e Engine) paintPlayerHealth(surf *sdl.Surface) {
	y := cfg.WinX / e.player.MaxLp * e.player.LP
	surf.FillRect(&sdl.Rect{X: 1, Y: 1, H: 20, W: cfg.WinX}, 0x03fcdb)
	surf.FillRect(&sdl.Rect{X: 1, Y: 1, H: 20, W: int32(y)}, 0xfc2403)
	text, err := e.font.RenderUTF8Blended(fmt.Sprintf("LP: %v/%v", e.player.LP, e.player.MaxLp), sdl.Color{R: 0, G: 0, B: 0, A: 255})
	if err != nil {
		return
	}
	defer text.Free()

	// Draw the text around the center of the window
	if err = text.Blit(nil, surf, &sdl.Rect{X: cfg.WinX/2 - text.W/2, Y: 0, W: 0, H: 0}); err != nil {
		return
	}
}

func (e Engine) paintStats(surf *sdl.Surface, text map[string]string) {
	keys := make([]string, 0, len(text))

	for k := range text {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	y := int32(25)
	for _, k := range keys {
		v := text[k]
		text, err := e.font.RenderUTF8Blended(fmt.Sprintf("%v: %v", k, v), sdl.Color{R: 0, G: 0, B: 0, A: 255})
		if err != nil {
			return
		}

		// Draw the text around the center of the window
		if err = text.Blit(nil, surf, &sdl.Rect{X: 2, Y: y, W: 0, H: 0}); err != nil {
			return
		}
		y += text.H + 2
		text.Free() // potential leek
	}

}
