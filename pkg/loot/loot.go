package loot

import "github.com/vogtp/eggsis-go/pkg/player"

type Loot func(*player.Player)

var NoLoot Loot = func(p *player.Player) {}