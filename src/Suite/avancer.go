package suite

import "fmt"

var villages = []string{
	"Séoul",
	"Busan",
	"Tokyo",
	"Île de Jeju",
}

func Menu(nom string, afficherInventaire func()) {
	position := 0

	for {
		fmt.Println("|--------------------------|")
		fmt.Printf(" Lieu actuel : %s\n", villages[position])
		fmt.Println("|--------------------------|")
		fmt.Println("Que veux-tu faire maintenant ?")
		fmt.Println("1 - Aller à la ville suivante")
		fmt.Println("2 - Ouvrir l'inventaire")
		fmt.Println("3 - Quitter")

		var choix string
		fmt.Scanln(&choix)

		switch choix {
		case "1":
			position = VillageSuivant(position)
		case "2":
			Inventaire(nom, afficherInventaire)
		case "3":
			Quitter(nom)
			return
		default:
			fmt.Println("Choix invalide, tapez 1, 2 ou 3.")
		}
	}
}
func VillageSuivant(position int) int {
	if position >= len(villages)-1 {
		fmt.Println("Tu es déjà à la dernière ville")
		return position
	}
	position++
	fmt.Println("Tu voyages vers :", villages[position])
	return position
}

// Inventaire affiche l'inventaire du joueur.
func Inventaire(nom string, afficherInventaire func()) {
	fmt.Println("Inventaire de", nom, ":")
	afficherInventaire()
}

// Quitter dit au revoir au joueur.
func Quitter(nom string) {
	fmt.Printf("À bientôt %s !\n", nom)
}
