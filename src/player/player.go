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
	HealthPerLevel = 15
	MaxLevel       = 50
)

type Player struct {
	Name       string
	Level      int
	Experience int
	Money      int
	MaxHealth  int
	Health     int
	Class      classes.Classe
	Inventory  Inventory
}

// MaxHealthForLevel calcule les PV max pour un niveau donné.
func MaxHealthForLevel(level int) int {
	if level < 1 {
		level = 1
	}

	return BaseMaxHealth + (level-1)*HealthPerLevel
}

func XPToNextLevel(level int) int {
	if level >= MaxLevel {
		return 0
	}

	return 50 * level
}

func initSkills(classe classes.Classe) []competence.Skill {
	skills := make([]competence.Skill, 0, len(classe.Competences))
	for _, comp := range classe.Competences {
		manaCost := 0
		if comp.Type == "magique" {
			manaCost = comp.Degats * 2
		}
		skills = append(skills, competence.NewSkill(
			comp.Nom,
			classe.Nom,
			comp.Type,
			manaCost,
			comp.Degats,
			1,
			comp.Type,
		))
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

	maxHP := MaxHealthForLevel(level)
	return Player{
		Name:       name,
		Level:      lv,
		Experience: 0,
		Money:      100,
		MaxHealth:  maxHP,
		Health:     maxHP,
		Class:      classe,
		Inventory:  NewInventory(),
		Skills:     initSkills(classe),
	}

	p.Weapon = weapon.NewArmeDeClasse(classe.Nom)
	p.Artefact = artefact.NouvelArtefactDeClasse(classe.Nom)

	if p.Artefact != nil {
		p.MaxHealth += p.Artefact.HealthBonus
		p.MaxMana += p.Artefact.ManaBonus
		p.Health = p.MaxHealth
		p.Mana = p.MaxMana
	}

	return p
}

func Character(nom string, classe classes.Classe) Player {
	return newPlayer(nom, 1, classe)
}

func NewLimule() Player {
	return newPlayer("Limule", 3, classes.NewMage())
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

	p.MaxHealth = MaxHealthForLevel(p.Level)
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

	p.Health -= damage

	if p.Health < 0 {
		p.Health = 0
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
		" Vie      : %d/%d\n",
		p.Health,
		p.MaxHealth,
	)

	fmt.Printf(
		" Mana     : %d/%d\n",
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

	if p.Weapon != nil {
		fmt.Printf(
			" Arme     : %s (+%d %s)\n",
			p.Weapon.Name,
			p.Weapon.AttackBonus,
			p.Weapon.Type,
		)
	}

	if p.Artefact != nil {
		fmt.Printf(
			" Artéfact : %s\n",
			p.Artefact.Name,
		)
	}

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
