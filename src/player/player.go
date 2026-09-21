package player

import (
	"fmt"

	classes "piscine/PlayerClass"
)

type Player struct {
	Name      string
	Level     int
	Money     int
	MaxHealth int
	Health    int
	Class     classes.Classe
	Inventory Inventory
}

func Character(nom string, classe classes.Classe) Player {
	return Player{
		Name:      nom,
		Level:     1,
		Money:     100,
		MaxHealth: 100,
		Health:    100,
		Class:     classe,
		Inventory: NewInventory(),
	}
}

func NewLimule() Player {
	return Player{
		Name:      "Limule",
		Level:     3,
		Money:     100,
		MaxHealth: 130,
		Health:    130,
		Class:     classes.NewMage(),
		Inventory: NewInventory(),
	}
}

func (p Player) Afficher() {
	fmt.Println("|--------------------------|")
	fmt.Printf(" Nom      : %s\n", p.Name)
	fmt.Printf(" Classe   : %s\n", p.Class.Nom)
	fmt.Printf(" Niveau   : %d\n", p.Level)
	fmt.Printf(" Vie      : %d/%d\n", p.Health, p.MaxHealth)
	fmt.Printf(" Argent   : %d\n", p.Money)
	fmt.Printf(" Mana     : %d\n", p.Class.Mana)
	fmt.Printf(" Attaque Physique : %d / Attaque Magique : %d\n", p.Class.AttaquePhysique, p.Class.AttaqueMagique)
	fmt.Println(" Compétences :")
	for _, c := range p.Class.Competences {
		fmt.Printf("   - %s (%d dégâts)\n", c.Nom, c.Degats)
	}
	fmt.Println("|--------------------------|")
}

func IsDead(p Player) bool {
	if p.Health <= 0 {
		fmt.Println("Tu est mort réaparition avec 50% de ta vie max")
		p.Health = p.MaxHealth / 2
		return true
	}
	return false
}
