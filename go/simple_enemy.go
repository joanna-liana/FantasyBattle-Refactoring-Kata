package codingdojo

import "math"

type SimpleEnemy struct {
	armor Armor
	buffs []Buff
}

func MakeSimpleEnemy(armor Armor, buffs []Buff) SimpleEnemy {
	return SimpleEnemy{armor, buffs}
}

func (e SimpleEnemy) GetBuffs() []Buff {
	return e.buffs
}

func (e SimpleEnemy) GetArmor() Armor {
	return e.armor
}

func (e SimpleEnemy) GetSoak() int32 {
	sum := 0.0

	for _, buff := range e.GetBuffs() {
		sum += buff.SoakModifier()
	}

	return int32(math.Round(float64(e.GetArmor().GetDamageSoak()) * (sum + 1.0)))
}
