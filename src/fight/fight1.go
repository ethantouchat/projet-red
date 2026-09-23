package fight

import (
	player "piscine/player"
)

const MANA_ATK_MAGIE = 5

type Goblin struct {
	Name       string
	Level      int
	MaxHealth  int
	Health     int
	Attack     int
	Experience int
	Gold       int
}

func NewGoblin(name string, level int, health int, attack int, experience int) *Goblin {
	gold := level*10 + experience/5
	return &Goblin{
		Name:       name,
		Level:      level,
		MaxHealth:  health,
		Health:     health,
		Attack:     attack,
		Experience: experience,
		Gold:       gold,
	}
}

func (g *Goblin) IsAlive() bool {
	return g.Health > 0
}

func (g *Goblin) TakeDamage(damage int) {
	g.Health -= damage

	if g.Health < 0 {
		g.Health = 0
	}
}

func CreerAllie(niveau int) *player.Player {
	allie := player.NewLimule()
	p := &allie
	p.SetLevel(niveau)
	return p
}
