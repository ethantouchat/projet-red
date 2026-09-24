package fight

func NewGoblin2() *Goblin {
	return &Goblin{
		Name:      "Goblin2",
		Level:     3,
		MaxHealth: 60,
		Health:    60,
		Attack:    10,
		Experience: 220,
		Gold:       47,
	}
}
func NewGoblin3() *Goblin {
	return &Goblin{
		Name:      "Goblin3",
		Level:     4,
		MaxHealth: 70,
		Health:    70,
		Attack:    15,
		Experience: 300,
		Gold:       64,
	}
}
func NewAdvancedGoblin() *Goblin {
	return &Goblin{
		Name:      "Advanced Goblin",
		Level:     3,
		MaxHealth: 150,
		Health:    150,
		Attack:    15,
		Experience: 500,
		Gold:       103,
		Drop:      "Goblin Helmet",
	}
}
