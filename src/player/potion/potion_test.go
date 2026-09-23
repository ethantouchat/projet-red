package potion

import (
	"testing"

	classes "piscine/PlayerClass"
	player "piscine/player"
)

func TestTakePotUsesHealingPotionWhenLowHealth(t *testing.T) {
	p := player.Character("Test", classes.NewMage())
	p.MaxHealth = 100
	p.Health = 10
	_ = p.Inventory.AddItem("potion de soin")
	_ = p.Inventory.AddItem("potion de poison")

	if !TakePot(&p) {
		t.Fatal("TakePot a retourné false alors qu'un soin était possible")
	}

	if p.Health != 60 {
		t.Fatalf("la vie attendue est 60, obtenue %d", p.Health)
	}

	if len(p.Inventory.Items) != 1 {
		t.Fatalf("une potion de soin doit avoir été consommée, inventaire : %d", len(p.Inventory.Items))
	}
}

func TestTakePotReturnsFalseWithoutPotion(t *testing.T) {
	p := player.Character("Test", classes.NewMage())
	p.MaxHealth = 100
	p.Health = 10

	if TakePot(&p) {
		t.Fatal("TakePot ne doit pas consommer de potion quand l'inventaire est vide")
	}

	if p.Health != 10 {
		t.Fatalf("la vie ne doit pas changer sans potion, obtenue %d", p.Health)
	}
}
