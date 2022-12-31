package codingdojo

import (
	"testing"
)

func TestDamageCalculations(t *testing.T) {
	t.Run("low soak", func(t *testing.T) {
		// given
		item := MakeBasicItem("test", 10, 1.2)
		inventory := MakeInventory(MakeEquipment(item, item, item, item, item))
		stats := MakeStats(0)

		target := MakeSimpleEnemy(MakeSimpleArmor(2), []Buff{MakeBasicBuff(2.20, 40)})
		player := MakePlayer(inventory, stats)

		// when
		damage := player.CalculateDamage(target)

		// then
		want := int32(294)
		got := damage.GetAmount()

		if want != got {
			t.Fatalf("Expected %v, got %v", want, got)
		}
	})

	t.Skip("high soak", func(t *testing.T) {
		// given
		item := MakeBasicItem("test", 10, 1.2)
		inventory := MakeInventory(MakeEquipment(item, item, item, item, item))
		stats := MakeStats(0)

		target := MakeSimpleEnemy(MakeSimpleArmor(200), []Buff{MakeBasicBuff(2.20, 40)})
		player := MakePlayer(inventory, stats)

		// when
		damage := player.CalculateDamage(target)

		// then
		want := int32(0)
		got := damage.GetAmount()

		if want != got {
			t.Fatalf("Expected %v, got %v", want, got)
		}
	})

	t.Run("player enemy", func(t *testing.T) {
		// given
		item := MakeBasicItem("test", 10, 1.2)
		inventory := MakeInventory(MakeEquipment(item, item, item, item, item))
		stats := MakeStats(0)

		target := MakePlayer(inventory, stats)
		player := MakePlayer(inventory, stats)

		// when
		damage := player.CalculateDamage(target)

		// then
		want := int32(0)
		got := damage.GetAmount()

		if want != got {
			t.Fatalf("Expected %v, got %v", want, got)
		}
	})

}
