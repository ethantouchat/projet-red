package fight

import (
	"piscine/artefact"
	"piscine/player"
 )

type Goblin struct {
	Name       string
	Level      int
	MaxHealth  int
	Health     int
	Attack     int
	Experience int
	Gold       int
	Drop       string
	Drops      []string
}

func NewGoblin(name string, level int, health int, attack int, experience int) *Goblin {
	return &Goblin{
		Name:       name,
		Level:      level,
		MaxHealth:  health,
		Health:     health,
		Attack:     attack,
		Experience: experience,
		Gold:       level * 10,
		Drops:      []string{"GoblinTeeth"},
	}
}

func (g *Goblin) IsAlive() bool {
	return g != nil && g.Health > 0
}

func (g *Goblin) TakeDamage(damage int) {
	g.Health -= damage
	if g.Health < 0 {
		g.Health = 0
	}
}

func CreerAllie(niveau int) *player.Player {
	allie := player.NewLimule()
	if allie.Artefact == nil {
		allie.EquiperArtefact(artefact.NewEquipment("Orb of Avarice"))
	}
	allie.SetLevel(niveau)
	return &allie
}
