package village

import (
	"fmt"
	"piscine/player"
)

// Marchand représente un commerçant de village.
type Marchand struct {
	Nom string
}

func NewMarchand() Marchand {
	return Marchand{Nom: "Le Marchand"}
}

// AcheterPotionSoin vend une potion de soin au joueur.
func (m Marchand) AcheterPotionSoin(j *player.Player) bool {
	if j == nil {
		return false
	}

	prix := 10
	if j.Money < prix {
		fmt.Println(m.Nom, ": Tu n'as pas assez d'argent pour une potion de soin.")
		return false
	}

	j.Money -= prix
	j.Inventory.AddItem("potion de soin")
	fmt.Println(m.Nom, ": Tu achètes une potion de soin pour 10 pièces.")
	return true
}

// AcheterPotionPoison vend une potion de poison au joueur.
func (m Marchand) AcheterPotionPoison(j *player.Player) bool {
	if j == nil {
		return false
	}

	prix := 15
	if j.Money < prix {
		fmt.Println(m.Nom, ": Tu n'as pas assez d'argent pour une potion de poison.")
		return false
	}

	j.Money -= prix
	j.Inventory.AddItem("potion de poison")
	fmt.Println(m.Nom, ": Tu achètes une potion de poison pour 15 pièces.")
	return true
}

// AcheterFer vend du fer au joueur pour fabriquer ou améliorer des objets.
func (m Marchand) AcheterFer(j *player.Player, materiaux *Materiaux) bool {
	if j == nil || materiaux == nil {
		return false
	}

	prix := 10
	if j.Money < prix {
		fmt.Println(m.Nom, ": Tu n'as pas assez d'argent pour acheter du fer.")
		return false
	}

	j.Money -= prix
	materiaux.Metal += 1
	fmt.Println(m.Nom, ": Tu achètes 1 fer pour 10 pièces.")
	return true
}

// Boutique affiche les objets vendus par le marchand.
func (m Marchand) Boutique(j *player.Player, materiaux *Materiaux) {
	fmt.Println("|--------------------------|")
	fmt.Println(m.Nom, ": Bienvenue, que veux-tu acheter ?")
	fmt.Println("1 - Potion de soin   - 10 pièces")
	fmt.Println("2 - Potion de poison - 15 pièces")
	fmt.Println("3 - Fer              - 10 pièces")
	fmt.Println("4 - Quitter")
	fmt.Println("|--------------------------|")

	var choix string
	fmt.Scanln(&choix)

	switch choix {
	case "1":
		m.AcheterPotionSoin(j)
	case "2":
		m.AcheterPotionPoison(j)
	case "3":
		m.AcheterFer(j, materiaux)
	case "4":
		fmt.Println(m.Nom, ": À bientôt !")
	default:
		fmt.Println(m.Nom, ": Choix invalide.")
	}
}
