package codingdojo

type Equipment struct {
	// TODO add a ring item that may be equipped
	//  that may also add damage modifier
	leftHand  Item
	rightHand Item
	head      Item
	feet      Item
	chest     Item
}

func MakeEquipment(leftHand Item, rightHand Item, head Item, feet Item, chest Item) Equipment {
	return Equipment{leftHand, rightHand, head, feet, chest}
}

func (e Equipment) GetBaseDamage() int32 {
	return e.leftHand.GetBaseDamage() +
		e.rightHand.GetBaseDamage() +
		e.head.GetBaseDamage() +
		e.feet.GetBaseDamage() +
		e.chest.GetBaseDamage()
}

func (e Equipment) GetDamageModifier() float64 {
	return e.leftHand.GetDamageModifier() +
		e.rightHand.GetDamageModifier() +
		e.head.GetDamageModifier() +
		e.feet.GetDamageModifier() +
		e.chest.GetDamageModifier()
}
