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
		fmt.Println("4 - Voir les équipements")
		fmt.Println("5 - Visiter le forgeron")
		fmt.Println("6 - Visiter le marchand")
		fmt.Println("7 - Faire un combat")
		fmt.Println("8 - Quitter")

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

			AfficherEquipements(p)

		case "5":

			village.VisiterForgeron(p)

		case "6":

			village.VisiterMarchand(p)

		case "7":

			ChoisirCombat(p, allies)

		case "8":

			Quitter(p.Name)
			return

		default:

			fmt.Println(
				"Choix invalide, tapez un nombre entre 1 et 8.",
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

func AfficherEquipements(p *player.Player) {
	fmt.Println()
	fmt.Println("|--------------------------|")
	fmt.Printf(" Équipements de %s\n", p.Name)
	fmt.Println("|--------------------------|")

	if p.Weapon == nil {
		fmt.Println("Aucune arme équipée.")
	} else {
		fmt.Printf("Arme : %s (rang %s)\n", p.Weapon.Name, p.Weapon.Rank)
		fmt.Printf("Bonus vie : +%d\n", p.Weapon.Life)
		fmt.Printf("Bonus mana : +%d\n", p.Weapon.Mana)
		fmt.Printf("Bonus attaque : +%d\n", p.Weapon.Attack)
	}

	if p.Artefact == nil {
		fmt.Println("Aucun artefact équipé.")
	} else {
		fmt.Printf("Artéfact : %s (rang %s)\n", p.Artefact.Name, p.Artefact.Rank)
		fmt.Printf("Bonus vie : +%d\n", p.Artefact.Life)
		fmt.Printf("Bonus mana : +%d\n", p.Artefact.Mana)
		fmt.Printf("Bonus attaque : +%d\n", p.Artefact.Attack)
		if p.Artefact.SpellDamageMultiplier > 1 {
			fmt.Printf("Dégâts des sorts : x%d\n", p.Artefact.SpellDamageMultiplier)
		}
		if p.Artefact.Name == "Orb of Avarice" {
			fmt.Println("Effet : Orbe d'Avarice, les dégâts des sorts sont doublés.")
		}
	}

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
		ennemis[1].Drops = append(ennemis[1].Drops, "Traditional Goblin Necklace")
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

	if p.ArtefactPasEquiper {
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
