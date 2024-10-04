package player

import "fmt"

func (p Player) Stats() map[string]string {
	stats := make(map[string]string)
	stats["LP"] = fmt.Sprintf("%v/%v", p.LP, p.MaxLp)
	stats["armor"] = fmt.Sprintf("%v", p.Armor)
	stats["Gold"] = fmt.Sprintf("%v",p.Gold)
	return stats
}
