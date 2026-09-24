package suite

import (
	"fmt"

	"piscine/fight"
	player "piscine/player"
	"piscine/village"
)

var villages = []string{
	"Séoul",
	"Busan",
	"Tokyo",
	"Île de Jeju",
}

func Menu(p *player.Player) {

	position := 0
	var allies []*player.Player

	for {

		fmt.Println()
		fmt.Println("|--------------------------|")
		fmt.Printf(" Lieu actuel : %s\n", villages[position])
		fmt.Println("|--------------------------|")

		fmt.Println("Que veux-tu faire maintenant ?")
		fmt.Println("1 - Aller à la ville suivante")
		fmt.Println("2 - Ouvrir l'inventaire")
		fmt.Println("3 - Voir mon personnage")
		fmt.Println("4 - Visiter le forgeron")
		fmt.Println("5 - Visiter le marchand")
		fmt.Println("6 - Faire un combat")
		fmt.Println("7 - Quitter")

		var choix string
		fmt.Scanln(&choix)

		switch choix {

		case "1":

			if position >= len(villages)-1 {
				fmt.Println("Tu es déjà à la dernière ville.")
				continue
			}

			victoire := Voyager(position, p, allies)

			if !victoire {
				fmt.Println()
				fmt.Println("Le voyage est interrompu.")
				continue
			}

			position++

			fmt.Println()
			fmt.Println("Tu arrives à :", villages[position])

			if position >= 1 && len(allies) == 0 {
				allie := fight.CreerAllie(p.Level)
				allies = []*player.Player{allie}
				fmt.Println()
				fmt.Printf("Un allié se joint à vous : %s (Epeiste, Niveau %d)\n", allie.Name, allie.Level)
			}

		case "2":

			Inventaire(p)

		case "3":

			p.Afficher()

		case "4":

			village.VisiterForgeron(p)

		case "5":

			village.VisiterMarchand(p)

		case "6":


			ChoisirCombat(p, allies)

		case "7":

			Quitter(p.Name)
			return

		default:

			fmt.Println(
				"Choix invalide, tapez un nombre entre 1 et 7.",
			)
		}
	}
}

func Inventaire(p *player.Player) {

	fmt.Println()
	fmt.Println("|--------------------------|")
	fmt.Printf(" Inventaire de %s\n", p.Name)
	fmt.Println("|--------------------------|")

	p.Inventory.Afficher()

	fmt.Println("|--------------------------|")
}

func Quitter(nom string) {

	fmt.Printf(
		"À bientôt %s !\n",
		nom,
	)
}

func ChoisirCombat(p *player.Player, allies []*player.Player) {

	fmt.Println()
	fmt.Println("Quel combat veux-tu lancer ?")
	fmt.Println("1 - Fight 1")
	fmt.Println("2 - Fight 2")

	var choix string
	fmt.Scanln(&choix)

	switch choix {

	case "1":

		ennemis := []*fight.Goblin{
			fight.NewGoblin("Goblin", 1, 20, 5, 25),
			fight.NewGoblin("Goblin", 1, 20, 5, 25),
		}
		fight.Combat(p, allies, ennemis)

	case "2":

			AvertirFight2(p)
		ennemis := []*fight.Goblin{
			fight.NewGoblin2(),
			fight.NewGoblin3(),
			fight.NewAdvancedGoblin(),
		}
		fight.Combat(p, allies, ennemis)

	default:

		fmt.Println("Choix invalide, tape 1 ou 2.")
	}
}

func AvertirFight2(p *player.Player) {

	avertissement := false

	if p.Level < 3 {
		fmt.Println("AVERTISSEMENT : Fight 2 est conseillé à partir du niveau 3.")
		avertissement = true
	}

	if !p.ArmeAmelioree {
		fmt.Println("AVERTISSEMENT : ton arme n'est pas encore améliorée.")
		avertissement = true
	}

	if !p.ArtefactPasEquiper {
		fmt.Println("AVERTISSEMENT : ton artefact n'est pas encore equiper.")
		avertissement = true
	}

	if avertissement {
		fmt.Println("Tu peux continuer, mais ce combat risque d'être difficile.")
	}
}

func Voyager(position int, p *player.Player, allies []*player.Player) bool {

	switch position {

	case 0:

		fmt.Println()
		fmt.Println("Tu quittes Séoul...")
		fmt.Println("Des Goblins apparaissent !")

		ennemis := []*fight.Goblin{
			fight.NewGoblin("Goblin", 1, 20, 5, 25),
			fight.NewGoblin("Goblin", 1, 20, 5, 25),
		}

		return fight.Combat(p, allies, ennemis)

	case 1:

		fmt.Println()
		fmt.Println("Tu quittes Busan...")
		fmt.Println("Un Goblin Elite apparaît !")

		ennemis := []*fight.Goblin{
			fight.NewGoblin("Goblin Elite", 2, 40, 8, 50),
		}

		return fight.Combat(p, allies, ennemis)

	case 2:

		fmt.Println()
		fmt.Println("Tu quittes Tokyo...")
		fmt.Println("Plusieurs ennemis apparaissent !")

		ennemis := []*fight.Goblin{
			fight.NewGoblin("Goblin", 2, 30, 7, 35),
			fight.NewGoblin("Goblin", 2, 30, 7, 35),
			fight.NewGoblin("Goblin Elite", 3, 60, 10, 80),
		}

		return fight.Combat(p, allies, ennemis)

	default:

		return false
	}

}
