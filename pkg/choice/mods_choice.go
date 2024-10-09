package choice

import (
	"time"
)

var Mods []Item

func init() {
	Mods = make([]Item, 0)

	Mods = append(Mods, damage("knife", 30, 10, "Do more damage with a knife"))
	Mods = append(Mods, damage("Sword", 60, 15, "Do lots of damage with a sword"))
	Mods = append(Mods, damage("Katana", 80, 25, "just a katana"))

	Mods = append(Mods, armor("Helmet", 50, 1, "Protect your egg with a helmet"))
	Mods = append(Mods, armor("Shield", 100, 2, "Protect your egg with a shield"))
	Mods = append(Mods, armor("Armor suit", 200, 5, "Protect your egg with suit of armor"))

	Mods = append(Mods, health("Corn", 10, 40, "Get some extra health by eating corn"))
	Mods = append(Mods, health("bug", 20, 80, "Get some extra health by eating a bug"))
	Mods = append(Mods, health("Worm", 40, 160, "Get some extra health by eating a worm"))

	Mods = append(Mods, health("Long legs", 20, 4, "Get some extra speed with long legs"))
	Mods = append(Mods, health("Wings", 40, 8, "Get some extra speed with wings"))
	Mods = append(Mods, health("Big Wings", 80, 16, "Get some extra speed with wings"))

	Mods = append(Mods, attackSpeed("Bar brawl", 30, 0.9, "Fight faster with bar brawling"))
	Mods = append(Mods, attackSpeed("Karate", 60, 0.8, "Fight faster with karate"))
	Mods = append(Mods, attackSpeed("Ninja", 120, 0.6, "Fight faster with ninja moves"))

	Mods = append(Mods, *newGun("Old gun", 20, 3, 20, 500*time.Millisecond, "Old slow gun").Item)
	Mods = append(Mods, *newGun("Good gun", 80, 10, 100, 100*time.Millisecond, "Strong fast gun gun").Item)
	Mods = append(Mods, *newGun("Machine gun", 200, 30, 120, 50*time.Millisecond, "Machine gun").Item)
}
