package codingdojo

type Inventory struct {
	equipment Equipment
}

func MakeInventory(equipment Equipment) Inventory {
	return Inventory{equipment}
}

func (i Inventory) GetBaseDamage() int32 {
	return i.equipment.GetBaseDamage()
}

func (i Inventory) GetDamageModifier() float64 {
	return i.equipment.GetDamageModifier()
}
