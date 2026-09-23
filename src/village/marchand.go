package village

import (
	"fmt"

	"piscine/artefact"
	player "piscine/player"
)

func VisiterMarchand(p *player.Player) {
	for {
		fmt.Println()
		fmt.Println("|--------------------------|")
		fmt.Println("       MARCHAND")
		fmt.Println("|--------------------------|")
		fmt.Println("1 - Acheter un artéfact")
		fmt.Println("2 - Acheter des potions")
		fmt.Println("3 - Quitter")

		var choix string
		fmt.Scanln(&choix)

		switch choix {
		case "1":
			acheterArtefact(p)
		case "2":
			acheterPotions(p)
		case "3":
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func acheterArtefact(p *player.Player) {
	artefacts := artefact.ListeArtefacts()

	fmt.Println()
	fmt.Println("|--------------------------|")
	fmt.Println(" Artéfacts disponibles :")
	for i, art := range artefacts {
		bonus := ""
		if art.HealthBonus > 0 {
			bonus += fmt.Sprintf("+%d PV ", art.HealthBonus)
		}
		if art.ManaBonus > 0 {
			bonus += fmt.Sprintf("+%d Mana ", art.ManaBonus)
		}
		if art.AttackBonus > 0 {
			bonus += fmt.Sprintf("+%d ATQ ", art.AttackBonus)
		}
		fmt.Printf(" %d - %s (%s)| %d or | %s\n",
			i+1, art.Name, bonus, art.Price, art.Rarity)
	}
	fmt.Println(" 0 - Annuler")
	fmt.Println("|--------------------------|")

	var choix int
	fmt.Scanln(&choix)

	if choix < 1 || choix > len(artefacts) {
		fmt.Println("Annulation.")
		return
	}

	art := artefacts[choix-1]

	if !p.RemoveMoney(art.Price) {
		fmt.Printf("Pas assez d'argent. Il vous manque %d or.\n", art.Price-p.Money)
		return
	}

	if p.Artefact != nil {
		fmt.Printf("Tu remplaces %s par %s.\n", p.Artefact.Name, art.Name)
	} else {
		fmt.Printf("Tu as acheté : %s.\n", art.Name)
	}

	p.Artefact = &art
	applyArtefactBonuses(p)
	fmt.Printf("Bonus : +%d PV, +%d Mana, +%d ATQ\n", art.HealthBonus, art.ManaBonus, art.AttackBonus)
}

func acheterPotions(p *player.Player) {
	fmt.Println()
	fmt.Println(" 1 - Potion de soin (50 or)")
	fmt.Println(" 2 - Potion de poison (80 or)")
	fmt.Println(" 0 - Annuler")

	var choix int
	fmt.Scanln(&choix)

	switch choix {
	case 1:
		if !p.RemoveMoney(50) {
			fmt.Println("Pas assez d'argent.")
			return
		}
		_ = p.Inventory.AddItem("potion de soin")
		fmt.Println("Tu as acheté une potion de soin.")
	case 2:
		if !p.RemoveMoney(80) {
			fmt.Println("Pas assez d'argent.")
			return
		}
		_ = p.Inventory.AddItem("potion de poison")
		fmt.Println("Tu as acheté une potion de poison.")
	case 0:
		fmt.Println("Annulation.")
	default:
		fmt.Println("Choix invalide.")
	}
}

func applyArtefactBonuses(p *player.Player) {
	if p.Artefact == nil {
		return
	}

	p.MaxHealth += p.Artefact.HealthBonus
	p.MaxMana += p.Artefact.ManaBonus
	if p.Health < p.MaxHealth {
		p.Health = p.MaxHealth
	}
	p.FullMana()
}
