package choice

import "github.com/vogtp/eggsis-go/pkg/player"

type PlayerModifier func(*player.Egg)

type Item struct {
	Name        string
	Image       string
	Description string
	Modifier    PlayerModifier
	Action      player.ActionFunc

	Cost int
}
