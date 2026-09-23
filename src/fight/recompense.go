package fight

import (
	"fmt"

	"piscine/player"
)

func RecompenseGoblin1(p *player.Player) {
	if p == nil {
		return
	}

	p.GainXP(150)
	_ = p.Inventory.AddItem("Goblin Teeth")
	_ = p.Inventory.AddItem("Goblin Teeth")
	_ = p.Inventory.AddItem("Iron")

	fmt.Println("Récompense Goblin 1 : +150 XP, 2 dents de goblin et du fer")
}

func RecompenseGoblin2(p *player.Player) {
	if p == nil {
		return
	}

	p.GainXP(220)
	_ = p.Inventory.AddItem("Goblin Teeth")
	_ = p.Inventory.AddItem("Leather")
	_ = p.Inventory.AddItem("Iron")

	fmt.Println("Récompense Goblin 2 : +220 XP, 1 dent de goblin, du cuir et du fer")
}

func RecompenseGoblin3(p *player.Player) {
	if p == nil {
		return
	}

	p.GainXP(300)
	_ = p.Inventory.AddItem("Goblin Teeth")
	_ = p.Inventory.AddItem("Goblin Teeth")
	_ = p.Inventory.AddItem("Steel")

	fmt.Println("Récompense Goblin 3 : +300 XP, 2 dents de goblin et de l'acier")
}

func RecompenseAdvancedGoblin(p *player.Player) {
	if p == nil {
		return
	}

	p.GainXP(500)
	_ = p.Inventory.AddItem("Goblin Teeth")
	_ = p.Inventory.AddItem("Goblin Helmet")
	_ = p.Inventory.AddItem("Steel")

	fmt.Println("Récompense Advanced Goblin : +500 XP, 1 dent de goblin, un casque et de l'acier")
}
