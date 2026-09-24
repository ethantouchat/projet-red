package fight

import (
	"fmt"
	"strings"

	"piscine/artefact"
	"piscine/player"
	"piscine/player/potion"
)

func Combat(p *player.Player, allies []*player.Player, ennemis []*Goblin) bool {
	if p == nil || len(ennemis) == 0 {
		return false
	}
	defer p.FullHeal()
	attaquesPhysicoMagiquesRestantes := 3

	fmt.Println("|--------------------------|")
	fmt.Println("          COMBAT")
	fmt.Println("|--------------------------|")

	for p.IsAlive() && ennemisVivants(ennemis) {
		potion.TakePot(p)
		fmt.Printf("%s : %d/%d PV | Mana : %d/%d\n", p.Name, p.Health, p.MaxHealth, p.Mana, p.MaxMana)
		afficherAllies(allies)
		afficherEnnemis(ennemis)
		fmt.Println("1 - Attaquer")
		fmt.Println("2 - Utiliser une compétence")
		fmt.Println("3 - Voir le personnage")
		fmt.Println("4 - Ouvrir l'inventaire")
		fmt.Println("5 - Fuir")
		if p.Class.Nom == "Assassin" {
			fmt.Printf("6 - Attaque physico-magique (%d restantes)\n", attaquesPhysicoMagiquesRestantes)
		}

		var choix string
		fmt.Scanln(&choix)

		switch choix {
		case "1":
			attaquer(p, ennemis)
		case "2":
			utiliserCompetence(p, allies, ennemis)
		case "3":
			p.Afficher()
		case "4":
			utiliserObjetCombat(p)
		case "5":
			fmt.Println("Tu prends la fuite !")
			return false
		case "6":
			attaquePhysicoMagique(p, ennemis, &attaquesPhysicoMagiquesRestantes)
		default:
			fmt.Println("Choix invalide.")
			continue
		}

		if ennemisVivants(ennemis) {
			tourAllies(allies, ennemis)
		}
		if ennemisVivants(ennemis) {
			tourEnnemis(p, ennemis)
			potion.TakePot(p)
		}
	}

	if !p.IsAlive() {
		fmt.Println("DEFAITE")
		return false
	}

	fmt.Println("VICTOIRE !")
	totalExperience := 0
	totalOr := 0
	recompenses := []string{}
	for _, ennemi := range ennemis {
		totalExperience += ennemi.Experience
		totalOr += ennemi.Gold
		p.GainXP(ennemi.Experience)
		p.AddMoney(ennemi.Gold)
		drops := append([]string{}, ennemi.Drops...)
		if ennemi.Drop != "" {
			drops = append(drops, ennemi.Drop)
		}
		for _, drop := range drops {
			if drop == "Traditional Goblin Necklace" && demanderEquipement(p, drop) {
				equipement := artefact.NewEquipment(drop)
				if p.EquiperArtefact(equipement) {
					recompenses = append(recompenses, drop+" équipé")
					continue
				}
			}
			if p.Inventory.AddItem(drop) == nil {
				recompenses = append(recompenses, drop)
			}
		}
	}
	fmt.Println("|--------------------------|")
	fmt.Println("       RÉCOMPENSES GAGNÉES")
	fmt.Printf("Expérience : +%d\n", totalExperience)
	fmt.Printf("Or : +%d pièces\n", totalOr)
	for _, recompense := range recompenses {
		fmt.Printf("Objet : %s\n", recompense)
	}
	fmt.Println("|--------------------------|")
	return true
}

func demanderEquipement(p *player.Player, nom string) bool {
	fmt.Printf("Tu as obtenu : %s. Veux-tu l'équiper ? (o/n)\n", nom)
	var choix string
	fmt.Scanln(&choix)
	return strings.EqualFold(choix, "o") || strings.EqualFold(choix, "oui")
}

func afficherEnnemis(ennemis []*Goblin) {
	fmt.Println("Ennemis :")
	for _, ennemi := range ennemis {
		if ennemi.IsAlive() {
			fmt.Printf("- %s : %d/%d PV\n", ennemi.Name, ennemi.Health, ennemi.MaxHealth)
		} else {
			fmt.Printf("- %s : MORT\n", ennemi.Name)
		}
	}
}

func afficherAllies(allies []*player.Player) {
	for _, allie := range allies {
		if allie == nil {
			continue
		}
		fmt.Printf("Allié %s : %d/%d PV | Mana : %d/%d\n", allie.Name, allie.Health, allie.MaxHealth, allie.Mana, allie.MaxMana)
	}
}

func ennemisVivants(ennemis []*Goblin) bool {
	for _, ennemi := range ennemis {
		if ennemi.IsAlive() {
			return true
		}
	}
	return false
}

func attaquer(p *player.Player, ennemis []*Goblin) {
	for _, ennemi := range ennemis {
		if ennemi.IsAlive() {
			damage := p.Class.AttaquePhysique
			ennemi.TakeDamage(damage)
			fmt.Printf("%s inflige %d dégâts à %s.\n", p.Name, damage, ennemi.Name)
			return
		}
	}
}

func attaquePhysicoMagique(p *player.Player, ennemis []*Goblin, restantes *int) {
	if p.Class.Nom != "Assassin" {
		fmt.Println("Cette attaque est réservée à l'Assassin.")
		return
	}
	if restantes == nil || *restantes <= 0 {
		fmt.Println("Tu n'as plus d'attaques physico-magiques pour ce combat.")
		return
	}

	for _, ennemi := range ennemis {
		if ennemi.IsAlive() {
			damage := p.Class.AttaquePhysique + p.Class.AttaqueMagique
			ennemi.TakeDamage(damage)
			*restantes--
			fmt.Printf("%s utilise son attaque physico-magique et inflige %d dégâts à %s.\n", p.Name, damage, ennemi.Name)
			return
			return
		}
	}
}

func utiliserCompetence(p *player.Player, allies []*player.Player, ennemis []*Goblin) {
	if len(p.Skills) == 0 {
		fmt.Println("Tu n'as aucune compétence.")
		return
	}

	fmt.Println("Choisis une compétence :")
	for i, skill := range p.Skills {
		fmt.Printf("%d - %s (niveau requis : %d, %d mana, +%d dégâts)\n", i+1, skill.Name, skill.RequiredLevel, skill.ManaCost, skill.AttackBonus)
	}

	var choix int
	fmt.Scanln(&choix)
	index := choix - 1
	if index < 0 || index >= len(p.Skills) {
		fmt.Println("Compétence invalide.")
		return
	}

	skill := p.Skills[index]
	if p.Level < skill.RequiredLevel {
		fmt.Printf("Niveau insuffisant : niveau %d requis.\n", skill.RequiredLevel)
		return
	}
	if p.Mana < skill.ManaCost {
		fmt.Println("Pas assez de mana.")
		return
	}

	var ennemi *Goblin
	for _, candidat := range ennemis {
		if candidat.IsAlive() {
			ennemi = candidat
			break
		}
	}
	if ennemi == nil {
		return
	}

	p.Mana -= skill.ManaCost
	if skill.Name == "Récupération de mana" {
		restaurerManaAllies(allies)
		return
	}

	damage := skill.AttackBonus
	if skill.Type == "magique" {
		damage += p.Class.AttaqueMagique
		if p.Artefact != nil {
			damage = artefact.ApplySpellDamage(damage, *p.Artefact)
		}
	} else {
		damage += p.Class.AttaquePhysique
	}

	ennemi.TakeDamage(damage)
	fmt.Printf("%s utilise %s et inflige %d dégâts à %s.\n", p.Name, skill.Name, damage, ennemi.Name)
}

func restaurerManaAllies(allies []*player.Player) {
	if len(allies) == 0 {
		fmt.Println("Aucun allié ne peut recevoir du mana.")
		return
	}

	for _, allie := range allies {
		if allie == nil || !allie.IsAlive() {
			continue
		}
		gain := allie.MaxMana / 2
		allie.RestoreMana(gain)
		fmt.Printf("%s récupère 50%% de mana : %d/%d.\n", allie.Name, allie.Mana, allie.MaxMana)
	}
}

func utiliserObjetCombat(p *player.Player) {
	fmt.Println("Inventaire de combat :")
	p.Inventory.Afficher()
	fmt.Println("1 - Utiliser une potion de mana")
	fmt.Println("2 - Retour")

	var choix string
	fmt.Scanln(&choix)
	if choix == "1" {
		potion.UseManaPotion(p)
	}
}

func tourEnnemis(p *player.Player, ennemis []*Goblin) {
	for _, ennemi := range ennemis {
		if ennemi.IsAlive() {
			p.TakeDamage(ennemi.Attack)
			fmt.Printf("%s inflige %d dégâts à %s.\n", ennemi.Name, ennemi.Attack, p.Name)
			return
		}
	}
}

func tourAllies(allies []*player.Player, ennemis []*Goblin) {
	for _, allie := range allies {
		if allie == nil || !allie.IsAlive() {
			continue
		}

		for _, ennemi := range ennemis {
			if !ennemi.IsAlive() {
				continue
			}

			damage := allie.Class.AttaquePhysique
			if allie.Mana >= 5 {
				allie.Mana -= 5
				damage = allie.Class.AttaqueMagique
				if allie.Artefact != nil {
					damage = artefact.ApplySpellDamage(damage, *allie.Artefact)
				}
				fmt.Printf("%s utilise une attaque magique et inflige %d dégâts à %s. Mana : %d/%d\n", allie.Name, damage, ennemi.Name, allie.Mana, allie.MaxMana)
			} else {
				fmt.Printf("%s inflige %d dégâts à %s avec une attaque physique.\n", allie.Name, damage, ennemi.Name)
			}
			ennemi.TakeDamage(damage)
			return
		}
	}
}
