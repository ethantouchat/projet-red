package piscine

func NewGoblin2() Goblin {
	return Goblin{
		Name:      "Goblin2",
		Level:     3,
		MaxHealth: 60,
		Health:    60,
		Attack:    10,
	}
}
func NewGoblin3() Goblin {
	return Goblin{
		Name:      "Goblin3",
		Level:     4,
		MaxHealth: 70,
		Health:    70,
		Attack:    15,
	}
}
func NewAdvancedGoblin() Goblin {
	return Goblin{
		Name:      "Advanced Goblin",
		Level:     3,
		MaxHealth: 150,
		Health:    150,
		Attack:    30,
	}
}
