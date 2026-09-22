package weapon

type Weapon struct {
	Name   string
	Rank   string
	Life   int
	Attack int
	Mana   int
}

func NewWeapon(name string) Weapon {
	switch name {
	case "Sword":
		return Weapon{Name: "Sword", Rank: "F", Attack: 1}
	case "Dual Daggers":
		return Weapon{Name: "Dual Daggers", Rank: "F", Attack: 2}
	case "Staff":
		return Weapon{Name: "Staff", Rank: "F", Mana: 1}
	case "Spellbook":
		return Weapon{Name: "Spellbook", Rank: "D", Mana: 5}
	case "Goblin Slayer Dual Daggers":
		return Weapon{Name: "Goblin Slayer Dual Daggers", Rank: "C", Attack: 8}
	case "Goblin Slayer Staff":
		return Weapon{Name: "Goblin Slayer Staff", Rank: "C", Mana: 10}
	case "Goblin Slayer Sword":
		return Weapon{Name: "Goblin Slayer Sword", Rank: "C", Attack: 10}
	default:
		return Weapon{}
	}
}
