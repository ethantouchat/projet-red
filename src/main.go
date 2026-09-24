package main

import (
	"fmt"

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
	fmt.Println("1 - Assassin : 5 Mana / 4 Attaque Physique / 3 Attaque Magique")
	fmt.Println("2 - Epeiste  : 1 Mana / 10 Attaque Physique / 1 Attaque Magique")
	fmt.Println("3 - Mage     : 10 Mana / 1 Attaque Physique / 10 Attaque Magique")
	fmt.Println("|--------------------------|")

	var classe classes.Classe
	for {
		fmt.Println("Veuillez choisir votre classe (1, 2 ou 3):")
		fmt.Scanln(&choix)

		nomClasse := map[string]string{
			"1": "assassin",
			"2": "epeiste",
			"3": "mage",
		}[choix]
		c, ok := classes.GetClasse(nomClasse)
		if ok {
			classe = c
			break
		}
		fmt.Println("Choix invalide, tape 1, 2 ou 3.")
	}
	fmt.Println("Vous avez choisi la classe", classe.Nom)

	joueur1 := player.Character(nom, classe)

	joueur1.Inventory.AddItem("potion de soin")
	joueur1.Inventory.AddItem("potion de poison")
	joueur1.Inventory.AddItem("potion de mana")

	joueur1.Afficher()

	suite.Menu(&joueur1)
}
