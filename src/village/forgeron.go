package village

import (
	"fmt"

	player "piscine/player"
	"piscine/weapon"
)

func VisiterForgeron(p *player.Player) {
	for {
		fmt.Println()
		fmt.Println("|--------------------------|")
		fmt.Println("       FORGERON")
		fmt.Println("|--------------------------|")
		fmt.Println("1 - Acheter une arme")
		fmt.Println("2 - Améliorer mon arme")
		fmt.Println("3 - Voir mon arme")
		fmt.Println("4 - Quitter")

		var choix string
		fmt.Scanln(&choix)

		switch choix {
		case "1":
			acheterArme(p)
		case "2":
			ameliorerArme(p)
		case "3":
			voirArme(p)
		case "4":
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func voirArme(p *player.Player) {
	fmt.Println()
	fmt.Println("|--------------------------|")
	if p.Weapon == nil {
		fmt.Println(" Tu n'as pas d'arme.")
	} else {
		fmt.Printf(" Arme : %s\n", p.Weapon.Name)
		fmt.Printf(" Bonus : +%d %s\n", p.Weapon.AttackBonus, p.Weapon.Type)
		fmt.Printf(" Niveau : %d\n", p.Weapon.Niveau)
		fmt.Printf(" Rareté : %s\n", p.Weapon.Rarity)
	}
	fmt.Println("|--------------------------|")
}

func acheterArme(p *player.Player) {
	armes := weapon.ListeArmes()

	fmt.Println()
	fmt.Println("|--------------------------|")
	fmt.Println(" Armes disponibles :")
	for i, arme := range armes {
		fmt.Printf(" %d - %s (+%d %s) | %d or | %s\n",
			i+1, arme.Name, arme.AttackBonus, arme.Type, arme.Price, arme.Rarity)
	}
	fmt.Println(" 0 - Annuler")
	fmt.Println("|--------------------------|")

	var choix int
	fmt.Scanln(&choix)

	if choix < 1 || choix > len(armes) {
		fmt.Println("Annulation.")
		return
	}

	arme := armes[choix-1]

	if !p.RemoveMoney(arme.Price) {
		fmt.Printf("Pas assez d'argent. Il vous manque %d or.\n", arme.Price-p.Money)
		return
	}

	p.Weapon = &arme
	fmt.Printf("Tu as acheté : %s (+%d %s)\n", arme.Name, arme.AttackBonus, arme.Type)
}

func ameliorerArme(p *player.Player) {
	if p.Weapon == nil {
		fmt.Println("Tu n'as pas d'arme à améliorer.")
		return
	}

	prix := p.Weapon.Niveau * 50
	if prix < 50 {
		prix = 50
	}

	fmt.Printf("Améliorer %s ? Coût : %d or + ", p.Weapon.Name, prix)

	matieres := []string{"Iron", "Steel"}
	matierePrincipale := matieres[p.Weapon.Niveau%2]

	fmt.Printf("1 %s\n", matierePrincipale)

	if p.Money < prix {
		fmt.Printf("Pas assez d'argent. Il vous manque %d or.\n", prix-p.Money)
		return
	}

	if p.Inventory.CountItem(matierePrincipale) < 1 {
		fmt.Printf("Pas assez de %s.\n", matierePrincipale)
		return
	}

	p.RemoveMoney(prix)
	p.Inventory.RemoveItem(matierePrincipale)
	weapon.AmeliorerArme(p.Weapon)

	fmt.Printf("%s améliorée au niveau %d ! (+%d %s)\n",
		p.Weapon.Name, p.Weapon.Niveau, p.Weapon.AttackBonus, p.Weapon.Type)
}
