package loot

import (
	"github.com/vogtp/eggsis-go/pkg/thing"
)

type Loot interface {
	Loot(*thing.Thing)
	Image() string
}
