package artefact

type Equipment struct {
	Name                  string
	Rank                  string
	Life                  int
	Attack                int
	Mana                  int
	SpellDamageMultiplier int
	AllowedPlayer string
}

func NewEquipment(nom string) Equipment {
	switch nom {
	case "Orb of Avarice":
		return Equipment{Name: "Orb of Avarice", Rank: "SS", SpellDamageMultiplier: 2, AllowedPlayer: "limule"}
	case "Goblin Ring":
		return Equipment{Name: "Goblin Ring", Rank: "B", Mana: 17, AllowedPlayer: "limule"}
	case "Traditional Goblin Necklace":
		return Equipment{Name: "Traditional Goblin Necklace", Rank: "C", Life: 20, Attack: 2, Mana: 2, AllowedPlayer: "Player"}
	case "Goblin Helmet":
		return Equipment{Name: "Goblin Helmet", Rank: "C", Life: 30, Attack: 5, AllowedPlayer: "Player"}
	default:	
		return Equipment{}
	}
}

func ApplySpellDamage(damage int, equipment Equipment) int {
	if equipment.SpellDamageMultiplier > 1 {
		return damage * equipment.SpellDamageMultiplier
	}

	return damage
}

func ApplyEquipmentBonuses(attack, health, mana int, helmetEquipped, necklaceEquipped, ringEquipped, orbEquipped bool) (int, int, int) {
	if helmetEquipped && necklaceEquipped {
		attack = attack * 110 / 100
		health = health * 110 / 100
	}

	if ringEquipped && orbEquipped {
		mana = mana * 80 / 100
		health = health * 120 / 100
	}

	return attack, health, mana
}
