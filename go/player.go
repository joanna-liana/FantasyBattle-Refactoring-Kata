package codingdojo

import "math"

type Player struct {
	inventory Inventory
	stats     Stats
}

func MakePlayer(inventory Inventory, stats Stats) Player {
	return Player{inventory, stats}
}

func (p Player) CalculateDamage(other Target) Damage {
	totalDamage := p.getTotalDamage()
	soak := other.GetSoak()

	return MakeDamage(max(0, totalDamage-soak))
}

func max(a int32, b int32) int32 {
	if a >= b {
		return a
	}
	return b
}

func (p Player) GetSoak() int32 {
	// TODO: Not implemented yet
	// Add friendly fire
	return p.getTotalDamage()
}

func (p Player) getTotalDamage() int32 {
	baseDamage := p.inventory.GetBaseDamage()
	damageModifier := p.stats.GetDamageModifier() + p.inventory.GetDamageModifier()

	return int32(math.Round(float64(baseDamage) * damageModifier))
}
