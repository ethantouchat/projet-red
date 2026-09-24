package player

import (
	"fmt"

	classes "piscine/PlayerClass"
	"piscine/artefact"
	"piscine/player/competence"
	"piscine/weapon"
)

const (
	BaseMaxHealth  = 100
	HealthPerLevel = 10
	MaxLevel       = 50
)

type Player struct {
	Name       string
	Level      int
	Experience int
	Money      int
	MaxHealth  int
	Health     int
	MaxMana    int
	Mana       int
	Class      classes.Classe
	Inventory  Inventory
	Skills     []competence.Skill
	Weapon     *weapon.Weapon
	Artefact   *artefact.Equipment
	ArmeAmelioree bool
	ArtefactPasEquiper bool
}

func XPToNextLevel(level int) int {
	if level >= MaxLevel {
		return 0
	}

	return 50 * level
}

func initSkills(classe classes.Classe, level int) []competence.Skill {
	skills := make([]competence.Skill, 0, len(classe.Competences))
	for _, comp := range classe.Competences {

		manaCost := comp.Degats * 2
		if manaCost == 0 {
			manaCost = 5
		}
		if classe.Nom == "Mage" && comp.Nom == "Boule de feu" {
			manaCost = 2
		}
		if classe.Nom == "Epeiste" {
			manaCost = 1
		}
		skill := competence.Skill{
			Name:          comp.Nom,
			Class:         classe.Nom,
			Type:          comp.Type,
			ManaCost:      manaCost,
			AttackBonus:   comp.Degats,
			RequiredLevel: 1,
			AllowedPlayer: "Player",
		}
		for skillLevel := 0; skillLevel < level; skillLevel++ {
			skill.LevelUp()
		}
		skills = append(skills, skill)
	}

	switch classe.Nom {
	case "Mage":
		skills = append(skills,
			competence.Skill{
				Name:          "Flame Emperor",
				Class:         "Mage",
				Type:          "magique",
				ManaCost:      25,
				AttackBonus:   35,
				RequiredLevel: 8,
			},
			competence.Skill{
				Name:          "Flash",
				Class:         "Mage",
				Type:          "magique",
				ManaCost:      10,
				AttackBonus:   10,
				RequiredLevel: 8,
			},
		)
	case "Assassin":
		skills = append(skills, competence.Skill{
			Name:          "Predator Aura",
			Class:         "Assassin",
			Type:          "physique",
			ManaCost:      10,
			AttackBonus:   10,
			RequiredLevel: 8,
		})
	}

	skills = append(skills, competence.Skill{
		Name:          "Récupération de mana",
		Class:         classe.Nom,
		Type:          "support",
		ManaCost:      2,
		RequiredLevel: 2,
	})

	for i := range skills {
		for skills[i].Level() < level {
			skills[i].LevelUp()
		}
	}

	return skills
}

func newPlayer(nom string, level int, classe classes.Classe) Player {

	if level < 1 {
		level = 1
	}

	if level > MaxLevel {
		level = MaxLevel
	}

	maxHP := BaseMaxHealth + (level-1)*HealthPerLevel
	nomArme := "Staff"
	switch classe.Nom {
	case "Assassin":
		nomArme = "Dual Daggers"
	case "Epeiste":
		nomArme = "Sword"
	}
	arme := weapon.NewWeapon(nomArme)
	p := Player{
		Name:       nom,
		Level:      level,
		Experience: 0,
		Money:      100,
		MaxHealth:  maxHP,
		Health:     maxHP,
		MaxMana:    classe.Mana,
		Mana:       classe.Mana,
		Class:      classe,
		Inventory:  NewInventory(),
		Skills:     initSkills(classe, level),
		Weapon:     &arme,
		ArtefactPasEquiper: true,
	}
	if classe.Nom == "Mage" {
		p.EquiperArtefact(artefact.NewEquipment("Orb of Avarice"))
	}

	return p
}

func (p *Player) EquiperArtefact(equipement artefact.Equipment) bool {
	if equipement.Name == "" || p.Artefact != nil {
		return false
	}

	p.Artefact = &equipement
	p.ArtefactPasEquiper = false
	p.MaxHealth += equipement.Life
	p.Health += equipement.Life
	p.MaxMana += equipement.Mana
	p.Mana += equipement.Mana
	p.Class.AttaquePhysique += equipement.Attack
	p.Class.AttaqueMagique += equipement.Attack

	return true
}

func Character(nom string, classe classes.Classe) Player {
	return newPlayer(nom, 1, classe)
}

func NewLimule() Player {
	limule := newPlayer("Limule", 3, classes.NewMage())
	limule.Class.AttaquePhysique = 10
	return limule
}

func (p *Player) SetLevel(level int) {
	if level < 1 {
		level = 1
	}
	if level > MaxLevel {
		level = MaxLevel
	}

	for p.Level < level {
		p.levelUp()
	}
}

func (p *Player) GainXP(amount int) {

	if amount <= 0 {
		return
	}

	if p.Level >= MaxLevel {
		return
	}

	p.Experience += amount

	fmt.Printf("+%d Experience\n", amount)

	for p.Level < MaxLevel {

		xpNecessaire := XPToNextLevel(p.Level)

		if p.Experience < xpNecessaire {
			break
		}

		p.Experience -= xpNecessaire

		p.levelUp()
	}
}

func (p *Player) levelUp() {

	if p.Level >= MaxLevel {
		return
	}

	p.Level++
	for i := range p.Skills {
		p.Skills[i].LevelUp()
		if p.Skills[i].RequiredLevel == p.Level {
			fmt.Printf("COMPÉTENCE DÉBLOQUÉE : %s\n", p.Skills[i].Name)
		}
	}

	p.MaxHealth = BaseMaxHealth + (p.Level-1)*HealthPerLevel
	p.MaxMana++
	p.Health = p.MaxHealth // soin complet à la montée de niveau

	fmt.Println("|--------------------------|")
	fmt.Printf(" NIVEAU SUPÉRIEUR ! Niveau %d\n", p.Level)
	fmt.Printf(" PV max : %d\n", p.MaxHealth)
	fmt.Printf(" Mana max : %d\n", p.MaxMana)
	fmt.Println("|--------------------------|")
}

func (p *Player) RestoreMana(amount int) {
	if amount <= 0 {
		return
	}

	p.Mana += amount

	if p.Mana > p.MaxMana {
		p.Mana = p.MaxMana
	}
}

func (p *Player) FullMana() {
	p.Mana = p.MaxMana
}

func (p *Player) FullHeal() {
	p.Health = p.MaxHealth
	p.Mana = p.MaxMana
}

func (p *Player) IsAlive() bool {
	return p.Health > 0
}

func (p *Player) IsDead() bool {

	if p.Health <= 0 {

		fmt.Println()
		fmt.Println("|--------------------------|")
		fmt.Println("          TU ES MORT")
		fmt.Println("|--------------------------|")

		fmt.Printf(
			"Réapparition avec 50%% de ta vie max : %d/%d PV\n",
			p.MaxHealth/2,
			p.MaxHealth,
		)

		p.Health = p.MaxHealth / 2

		return true
	}

	return false
}

func (p *Player) TakeDamage(damage int) {

	if damage <= 0 {
		return
	}

	wasAlive := p.Health > 0
	p.Health -= damage

	if p.Health < 0 {
		p.Health = 0
	}

	if wasAlive && p.Health == 0 {
		fmt.Println()
		fmt.Println("========================================")
		fmt.Println("              YOUR DEAD")
		fmt.Println("========================================")
	}
}

func (p *Player) Heal(amount int) {

	if amount <= 0 {
		return
	}

	p.Health += amount

	if p.Health > p.MaxHealth {
		p.Health = p.MaxHealth
	}
}

func (p *Player) AddMoney(amount int) {

	if amount <= 0 {
		return
	}

	p.Money += amount
}

func (p *Player) RemoveMoney(amount int) bool {

	if amount <= 0 {
		return true
	}

	if p.Money < amount {
		return false
	}

	p.Money -= amount

	return true
}

func (p Player) Afficher() {

	fmt.Println("|--------------------------|")
	fmt.Printf(" Nom      : %s\n", p.Name)
	fmt.Printf(" Classe   : %s\n", p.Class.Nom)

	if p.Level >= MaxLevel {
		fmt.Printf(
			" Niveau   : %d (NIVEAU MAX)\n",
			p.Level,
		)
	} else {
		fmt.Printf(
			" Niveau   : %d (Experience : %d/%d)\n",
			p.Level,
			p.Experience,
			XPToNextLevel(p.Level),
		)
	}

	fmt.Printf(
		" Vie      : %d/%d | Mana : %d/%d\n",
		p.Health,
		p.MaxHealth,
		p.Mana,
		p.MaxMana,
	)

	fmt.Printf(
		" Argent   : %d\n",
		p.Money,
	)

	fmt.Printf(
		" Attaque Physique : %d / Attaque Magique : %d\n",
		p.Class.AttaquePhysique,
		p.Class.AttaqueMagique,
	)
	fmt.Println(" Compétences :")

	for _, s := range p.Skills {
		fmt.Printf(
			"   - %s | Niv %d | +%d dégâts | Type: %s | Coût: %d mana\n",
			s.Name,
			s.Level(),
			s.AttackBonus,
			s.Type,
			s.ManaCost,
		)
	}

	fmt.Println("|--------------------------|")
}
