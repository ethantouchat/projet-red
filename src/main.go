package main

import (
	"fmt"
	"strings"

	classes "piscine/PlayerClass"
	suite "piscine/Suite"
	"piscine/player"
)

func main() {
	var nom string
	var choix string

	fmt.Println("|--------------------------|")
	fmt.Println("         Bienvenue")
	fmt.Println("|--------------------------|")

	fmt.Println("Création de votre personnage")
	fmt.Println("Veuillez choisir votre nom:")
	fmt.Scanln(&nom)
	fmt.Printf("Votre personnage s'appelle maintenant: %s\n", nom)

	fmt.Println("|--------------------------|")
	fmt.Println("    Information Classes")
	fmt.Println(" - Assassin: 5 Mana / 4 Attaque Physique / 3 Attaque Magique")
	fmt.Println(" - Mage: 10 Mana / 1 Attaque Physique / 10 Attaque Magique")
	fmt.Println(" - Epeiste: 1 Mana / 10 Attaque Physique / 1 Attaque Magique")
	fmt.Println("|--------------------------|")

	// On redemande tant que la classe n'est pas valide
	var classe classes.Classe
	for {
		fmt.Println("Veuillez choisir votre classe:")
		fmt.Scanln(&choix)
		choix = strings.ToLower(choix)

		c, ok := classes.GetClasse(choix)
		if ok {
			classe = c
			break
		}
		fmt.Println("Classe non reconnue, veuillez choisir entre Assassin, Epeiste et Mage")
	}
	fmt.Println("Vous avez choisi la classe", classe.Nom)

	joueur1 := player.Character(nom, classe)

	joueur1.Inventory.AddItem("potion de soin")
	joueur1.Inventory.AddItem("potion de poison")

	joueur1.Afficher()

	// Menu principal : boucle jusqu'à ce que le joueur quitte
	suite.Menu(joueur1.Name, joueur1.Inventory.Afficher)
}
