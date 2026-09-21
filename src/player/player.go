package player

import (
	"fmt"

	classes "piscine/PlayerClass"
)

type Player struct {
	Nom         string
	Niveau      int
	Argent      int
	VieMax      int
	VieActuelle int
	Classe      classes.Classe
	Inventaire  Inventory
}

func NewPlayer(nom string, classe classes.Classe) Player {
	return Player{
		Nom:         nom,
		Niveau:      1,
		Argent:      100,
		VieMax:      100,
		VieActuelle: 100,
		Classe:      classe,
		Inventaire:  NewInventory(),
	}
}

func NewLimule() Player {
	return Player{
		Nom:         "Limule",
		Niveau:      3,
		Argent:      100,
		VieMax:      130,
		VieActuelle: 130,
		Classe:      classes.NewMage(),
		Inventaire:  NewInventory(),
	}
}

func (p Player) Afficher() {
	fmt.Println("|--------------------------|")
	fmt.Printf(" Nom      : %s\n", p.Nom)
	fmt.Printf(" Classe   : %s\n", p.Classe.Nom)
	fmt.Printf(" Niveau   : %d\n", p.Niveau)
	fmt.Printf(" Vie      : %d/%d\n", p.VieActuelle, p.VieMax)
	fmt.Printf(" Argent   : %d\n", p.Argent)
	fmt.Printf(" Mana     : %d\n", p.Classe.Mana)
	fmt.Printf(" Attaque Physique : %d / Attaque Magique : %d\n", p.Classe.AttaquePhysique, p.Classe.AttaqueMagique)
	fmt.Println(" Compétences :")
	for _, c := range p.Classe.Competences {
		fmt.Printf("   - %s (%d dégâts)\n", c.Nom, c.Degats)
	}
	fmt.Println("|--------------------------|")
}
