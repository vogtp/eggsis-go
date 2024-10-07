package loot

import "github.com/vogtp/eggsis-go/pkg/player"

type Loot interface {
	Loot(*player.Egg)
	Image() string
}
