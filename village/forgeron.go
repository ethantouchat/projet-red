package village

import "fmt"

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
