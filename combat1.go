package piscine

type Goblin struct {
	Name      string
	Level     int
	MaxHealth int
	Health    int
	Attack    int
}

func initGoblin() Goblin {
	return Goblin{
		Name:      "Goblin1",
		Level:     1,
		MaxHealth: 5,
		Health:    5,
		Attack:    5,
	}
}
