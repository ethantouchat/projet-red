package fight

import (
	"fmt"

	player "piscine/player"
	"piscine/player/potion"
)

func Combat(p *player.Player, allies []*player.Player, ennemis []*Goblin) bool {

	fmt.Println()
	fmt.Println("|--------------------------|")
	fmt.Println("          COMBAT")
	fmt.Println("|--------------------------|")

	fmt.Printf(
		"%s rencontre %d ennemi(s) !\n",
		p.Name,
		len(ennemis),
	)

	for p.IsAlive() && ennemisVivants(ennemis) {

		potion.TakePot(p)

		for _, allie := range allies {
			if allie != nil && allie.IsAlive() {
				potion.TakePot(allie)
			}
		}

		afficherCombat(p, allies, ennemis)

		fmt.Println()
		fmt.Println("Que veux-tu faire ?")
		fmt.Println("1 - Attaquer")
		fmt.Println("2 - Voir les compétences")
		fmt.Println("3 - Utiliser potion de poison")
		fmt.Println("4 - Fuir")

		var choix string
		fmt.Scanln(&choix)

		switch choix {

		case "1":

			attaquer(p, ennemis)

			if ennemisVivants(ennemis) {
				tourAlliés(allies, ennemis)
				if ennemisVivants(ennemis) {
					tourEnnemis(p, allies, ennemis)
				}
			}

		case "2":

			afficherCompetences(p)

		case "3":

			utiliserPotionPoison(p, ennemis)

		case "4":

			fmt.Println()
			fmt.Println("Tu prends la fuite !")
			p.FullMana()
			return false

		default:

			fmt.Println("Choix invalide.")
		}
	}

	if !ennemisVivants(ennemis) && p.IsAlive() {

		fmt.Println()
		fmt.Println("|--------------------------|")
		fmt.Println("         VICTOIRE !")
		fmt.Println("|--------------------------|")

		gagnerExperience(p, allies, ennemis)

		p.FullMana()
		for _, allie := range allies {
			if allie != nil && allie.IsAlive() {
				allie.FullMana()
			}
		}

		return true
	}

	if !p.IsAlive() {

		fmt.Println()
		fmt.Println("|--------------------------|")
		fmt.Println("         DEFAITE")
		fmt.Println("|--------------------------|")

		p.IsDead()

		return false
	}

	return false
}

func armeBonus(p *player.Player, attackType string) int {
	if p.Weapon == nil || p.Weapon.Type != attackType {
		return 0
	}
	return p.Weapon.AttackBonus
}

func afficherCombat(p *player.Player, allies []*player.Player, ennemis []*Goblin) {

	fmt.Println()
	fmt.Println("|--------------------------|")

	fmt.Printf(
		" %s | Niv %d | PV %d/%d | Mana %d/%d\n",
		p.Name,
		p.Level,
		p.Health,
		p.MaxHealth,
		p.Mana,
		p.MaxMana,
	)

	if p.Weapon != nil {
		fmt.Printf(" Arme : %s (+%d %s)\n", p.Weapon.Name, p.Weapon.AttackBonus, p.Weapon.Type)
	}

	fmt.Println("|--------------------------|")

	for _, allie := range allies {
		if allie == nil {
			continue
		}
		if allie.IsAlive() {
			fmt.Printf(
				" %s | Niv %d | PV %d/%d | Mana %d/%d\n",
				allie.Name,
				allie.Level,
				allie.Health,
				allie.MaxHealth,
				allie.Mana,
				allie.MaxMana,
			)
		} else {
			fmt.Printf(" %s | [MORT]\n", allie.Name)
		}
	}

	if len(allies) > 0 {
		fmt.Println("|--------------------------|")
	}

	for i, ennemi := range ennemis {

		if !ennemi.IsAlive() {
			fmt.Printf(" %d - %s [MORT]\n", i+1, ennemi.Name)
			continue
		}

		fmt.Printf(
			" %d - %s | Niv %d | PV %d/%d | ATK %d\n",
			i+1,
			ennemi.Name,
			ennemi.Level,
			ennemi.Health,
			ennemi.MaxHealth,
			ennemi.Attack,
		)
	}

	fmt.Println("|--------------------------|")
}

func attaquer(p *player.Player, ennemis []*Goblin) {

	fmt.Println()
	fmt.Println("Quel ennemi veux-tu attaquer ?")

	aliveIndices := []int{}
	for i, ennemi := range ennemis {
		if !ennemi.IsAlive() {
			continue
		}
		aliveIndices = append(aliveIndices, i)
		fmt.Printf(
			"%d - %s (%d/%d PV)\n",
			len(aliveIndices),
			ennemi.Name,
			ennemi.Health,
			ennemi.MaxHealth,
		)
	}

	if len(aliveIndices) == 0 {
		fmt.Println("Aucun ennemi vivant.")
		return
	}

	var choix int
	fmt.Scanln(&choix)

	index := choix - 1

	if index < 0 || index >= len(aliveIndices) {
		fmt.Println("Ennemi invalide.")
		return
	}

	ennemi := ennemis[aliveIndices[index]]

	if !ennemi.IsAlive() {
		fmt.Println("Cet ennemi est déjà mort.")
		return
	}

	fmt.Println()
	fmt.Println("Choisis ton action :")
	fmt.Println("1 - Attaque de base (physique)")
	fmt.Printf("2 - Attaque magique (%d mana)\n", MANA_ATK_MAGIE)

	skillCount := 0
	for _, skill := range p.Skills {
		if skill.AttackBonus > 0 || skill.ManaCost > 0 {
			skillCount++
		}
	}

	for i, skill := range p.Skills {
		if skill.AttackBonus > 0 || skill.ManaCost > 0 {
			fmt.Printf(
				"%d - %s (dégâts: %d, type: %s, coût: %d mana)\n",
				i+3,
				skill.Name,
				skill.AttackBonus,
				skill.Type,
				skill.ManaCost,
			)
		}
	}

	if skillCount == 0 {
		fmt.Println("3 - (aucune compétence disponible)")
	}

	var action int
	fmt.Scanln(&action)

	degats := 0
	nomAttaque := ""

	if action == 1 {
		degats = p.Class.AttaquePhysique + armeBonus(p, "physique")
		nomAttaque = "Attaque de base"
	} else if action == 2 {
		if p.Mana < MANA_ATK_MAGIE {
			fmt.Printf("Pas assez de mana. (disponible: %d/%d)\n", p.Mana, p.MaxMana)
			return
		}
		p.Mana -= MANA_ATK_MAGIE
		degats = p.Class.AttaqueMagique + armeBonus(p, "magique")
		nomAttaque = "Attaque magique"
		fmt.Printf("Tu utilises %d mana. Mana restant : %d/%d\n", MANA_ATK_MAGIE, p.Mana, p.MaxMana)
	} else {

		skillIndex := action - 3
		if skillIndex < 0 || skillIndex >= len(p.Skills) {
			fmt.Println("Action invalide.")
			return
		}

		skill := &p.Skills[skillIndex]

		if skill.AttackBonus <= 0 {
			fmt.Println("Cette compétence ne fait pas de dégâts.")
			return
		}

		if skill.ManaCost > p.Mana {
			fmt.Printf("Pas assez de mana. (disponible: %d/%d)\n", p.Mana, p.MaxMana)
			return
		}

		if skill.ManaCost > 0 {
			p.Mana -= skill.ManaCost
			fmt.Printf("Tu utilises %d mana. Mana restant : %d/%d\n", skill.ManaCost, p.Mana, p.MaxMana)
		}

		nomAttaque = skill.Name
		if skill.Type == "magique" {
			degats = skill.AttackBonus + (p.Class.AttaqueMagique+armeBonus(p, "magique"))/3
		} else {
			degats = skill.AttackBonus + (p.Class.AttaquePhysique+armeBonus(p, "physique"))/3
		}

		if skill.Type == "magique" && skill.AttackBonus > 0 {
			if skill.GainXP(10) {
				fmt.Printf("%s gagne un niveau ! (Niveau %d)\n", skill.Name, skill.Level())
			} else {
				fmt.Printf("%s gagne de l'XP ! (Niveau %d)\n", skill.Name, skill.Level())
			}
		}
	}

	ennemi.TakeDamage(degats)

	fmt.Printf(
		"%s utilise %s sur %s et inflige %d dégâts !\n",
		p.Name,
		nomAttaque,
		ennemi.Name,
		degats,
	)

	if !ennemi.IsAlive() {
		fmt.Printf("%s est mort !\n", ennemi.Name)
	}
}

func tourAlliés(allies []*player.Player, ennemis []*Goblin) {

	for _, allie := range allies {
		if allie == nil || !allie.IsAlive() || !ennemisVivants(ennemis) {
			continue
		}

		ennemiCible := ciblePrioritaire(ennemis)
		if ennemiCible == nil {
			continue
		}

		degats := allie.Class.AttaquePhysique
		if allie.Weapon != nil && allie.Weapon.Type == "physique" {
			degats += allie.Weapon.AttackBonus
		}

		ennemiCible.TakeDamage(degats)

		fmt.Printf(
			"%s attaque %s et inflige %d dégâts !\n",
			allie.Name,
			ennemiCible.Name,
			degats,
		)

		if !ennemiCible.IsAlive() {
			fmt.Printf("%s est mort !\n", ennemiCible.Name)
		}
	}
}

func ciblePrioritaire(ennemis []*Goblin) *Goblin {
	var cible *Goblin
	for _, ennemi := range ennemis {
		if !ennemi.IsAlive() {
			continue
		}
		if cible == nil || ennemi.Health > cible.Health {
			cible = ennemi
		}
	}
	return cible
}

func tourEnnemis(p *player.Player, allies []*player.Player, ennemis []*Goblin) {

	fmt.Println()
	fmt.Println("--- Tour des ennemis ---")

	for _, ennemi := range ennemis {

		if !ennemi.IsAlive() {
			continue
		}

		p.TakeDamage(ennemi.Attack)

		fmt.Printf(
			"%s attaque %s et inflige %d dégâts !\n",
			ennemi.Name,
			p.Name,
			ennemi.Attack,
		)

		fmt.Printf("PV restants : %d/%d\n", p.Health, p.MaxHealth)

		for _, allie := range allies {
			if allie == nil || !allie.IsAlive() {
				continue
			}

			allie.TakeDamage(ennemi.Attack)

			fmt.Printf(
				"%s attaque %s et inflige %d dégâts !\n",
				ennemi.Name,
				allie.Name,
				ennemi.Attack,
			)

			fmt.Printf("PV restants de %s : %d/%d\n", allie.Name, allie.Health, allie.MaxHealth)

			if !allie.IsAlive() {
				fmt.Printf("%s est mort !\n", allie.Name)
			}
		}

		if !p.IsAlive() {
			return
		}
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

func afficherCompetences(p *player.Player) {

	fmt.Println()
	fmt.Println("|--------------------------|")
	fmt.Println("       COMPETENCES")
	fmt.Println("|--------------------------|")

	if len(p.Skills) == 0 {
		fmt.Println("Aucune compétence.")
		fmt.Println("|--------------------------|")
		return
	}

	for i, skill := range p.Skills {
		fmt.Printf(
			" %d - %s | +%d dégâts | %s | coût: %d mana | Niv: %d\n",
			i+1,
			skill.Name,
			skill.AttackBonus,
			skill.Type,
			skill.ManaCost,
			skill.Level(),
		)
	}

	fmt.Println("|--------------------------|")
}

func utiliserPotionPoison(p *player.Player, ennemis []*Goblin) {

	aliveIndices := []int{}
	for i, ennemi := range ennemis {
		if !ennemi.IsAlive() {
			continue
		}
		aliveIndices = append(aliveIndices, i)
		fmt.Printf(
			"%d - %s (%d/%d PV)\n",
			len(aliveIndices),
			ennemi.Name,
			ennemi.Health,
			ennemi.MaxHealth,
		)
	}

	if len(aliveIndices) == 0 {
		fmt.Println("Aucun ennemi vivant.")
		return
	}

	fmt.Println("Sur quel ennemi utiliser la potion de poison ?")

	var choix int
	fmt.Scanln(&choix)

	index := choix - 1

	if index < 0 || index >= len(aliveIndices) {
		fmt.Println("Cible invalide.")
		return
	}

	target := ennemis[aliveIndices[index]]

	potion.UsePoisonPotion(p, &target.Health, target.MaxHealth)
}

func gagnerExperience(p *player.Player, allies []*player.Player, ennemis []*Goblin) {

	totalXP := 0
	totalGold := 0

	for _, ennemi := range ennemis {
		totalXP += ennemi.Experience
		totalGold += ennemi.Gold
	}

	if totalXP <= 0 && totalGold <= 0 {
		return
	}

	if len(allies) == 0 {
		p.GainXP(totalXP)
		p.AddMoney(totalGold)
		if totalGold > 0 {
			fmt.Printf("Tu as gagné %d or !\n", totalGold)
		}
		return
	}

	xpJoueur := totalXP * 2 / 3
	xpAllie := totalXP / (3 * len(allies))

	p.GainXP(xpJoueur)
	p.AddMoney(totalGold)
	if totalGold > 0 {
		fmt.Printf("Tu as gagné %d or !\n", totalGold)
	}

	for _, allie := range allies {
		if allie != nil && allie.IsAlive() {
			allie.GainXP(xpAllie)
		}
	}
}
