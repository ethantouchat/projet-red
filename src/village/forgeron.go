package village

import (
	"fmt"
	"strings"
	"piscine/player"
	"piscine/weapon"
)

// Materiaux représente ce qu'il faut apporter au forgeron pour améliorer une arme.
type Materiaux struct {
	Metal       int
	GoblinTeeth int
}

// Forgeron représente le personnage du village qui améliore les armes.
type Forgeron struct {
	Nom string
}

func NewForgeron() Forgeron {
	return Forgeron{Nom: "Le Forgeron"}
}

// AmeliorerArme vérifie si le joueur a assez d'argent et de matériaux.
// Si oui, l'arme est améliorée. Sinon, le forgeron refuse et demande de revenir avec ce qu'il faut.
func (f Forgeron) AmeliorerArme(argent *int, materiaux *Materiaux, arme *string) bool {
	if argent == nil || materiaux == nil || arme == nil {
		fmt.Println(f.Nom, ": Tu me donnes quoi exactement ?")
		return false
	}

	cout := 50
	if *argent < cout {
		fmt.Println(f.Nom, ": Tu n'as pas assez d'argent. Il me faut 50 pièces.")
		fmt.Println(f.Nom, ": Va chercher des matériaux et reviens.")
		return false
	}

	if materiaux.GoblinTeeth < 2 || materiaux.Metal < 1 {
		fmt.Println(f.Nom, ": Tu manques encore des matériaux.")
		fmt.Println(f.Nom, ": Il me faut 2 goblinTeeth et 1 métal.")
		return false
	}

	*argent -= cout
	materiaux.Metal -= 1
	materiaux.GoblinTeeth -= 2

	switch *arme {
	case "Sword":
		*arme = "Goblin Slayer Sword"
	case "Staff":
		*arme = "Goblin Slayer Staff"
	case "Dual Daggers":
		*arme = "Goblin Slayer Dual Daggers"
	default:
		*arme = *arme + " +1"
	}

	fmt.Println("|--------------------------|")
	fmt.Println(f.Nom, ": Très bien. J'améliore ton arme.")
	fmt.Printf("Ton arme devient : %s\n", *arme)
	fmt.Println("|--------------------------|")
	return true
}

// DemanderMatieres affiche la liste des matériaux demandés.
func (f Forgeron) DemanderMatieres() {
	fmt.Println("|--------------------------|")
	fmt.Println(f.Nom, ": Pour améliorer ton arme, il me faut :")
	fmt.Println("- 50 pièces d'or")
	fmt.Println("- 1 métal")
	fmt.Println("- 2 goblinTeeth")
	fmt.Println("|--------------------------|")
}

func VisiterForgeron(j *player.Player) {
	if j == nil {
		return
	}

	forgeron := NewForgeron()
	if j.ArmeAmelioree {
		fmt.Println(forgeron.Nom, ": Ton arme est déjà améliorée.")
		return
	}

	if j.Money < 50 || j.Inventory.CountItem("GoblinTeeth") < 2 || j.Inventory.CountItem("Metal") < 1 {
		forgeron.DemanderMatieres()
		return
	}

	fmt.Println(forgeron.Nom, ": Tu as tous les matériaux nécessaires pour améliorer ton arme.")
	fmt.Println("Veux-tu améliorer ton arme ? (o/n)")
	var choix string
	fmt.Scanln(&choix)
	if !strings.EqualFold(choix, "o") && !strings.EqualFold(choix, "oui") {
		fmt.Println(forgeron.Nom, ": D'accord, ton arme reste inchangée.")
		return
	}

	j.Money -= 50
	j.Inventory.RemoveItems("GoblinTeeth", 2)
	j.Inventory.RemoveItem("Metal")
	ameliorerArme(j)
	j.ArmeAmelioree = true
	fmt.Println(forgeron.Nom, ": Ton arme est maintenant améliorée !")
}

func ameliorerArme(j *player.Player) {
	if j.Weapon == nil {
		return
	}

	nomAmeliore := map[string]string{
		"Sword":        "Goblin Slayer Sword",
		"Staff":        "Goblin Slayer Staff",
		"Dual Daggers": "Goblin Slayer Dual Daggers",
	}[j.Weapon.Name]
	if nomAmeliore == "" {
		return
	}

	ancienneMana := j.Weapon.Mana
	nouvelleArme := weapon.NewWeapon(nomAmeliore)
	*j.Weapon = nouvelleArme
	bonusMana := nouvelleArme.Mana - ancienneMana
	j.MaxMana += bonusMana
	j.Mana += bonusMana
}
