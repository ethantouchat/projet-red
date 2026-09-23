package player

import (
	"fmt"

	classes "piscine/PlayerClass"
)

const (
	BaseMaxHealth  = 100 // PV max au niveau 1
	HealthPerLevel = 15  // PV max gagnés par niveau (130-100)/(3-1) = 15
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
	return BaseMaxHealth + (level-1)*HealthPerLevel
}

// XPToNextLevel donne l'XP nécessaire pour passer au niveau suivant.
func XPToNextLevel(level int) int {
	return 50 * level // niv 1->2 : 50, niv 2->3 : 100, ...
}

// newPlayer crée un joueur à un niveau donné, avec des stats cohérentes.
func newPlayer(nom string, level int, classe classes.Classe) Player {
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
	}
}

func Character(nom string, classe classes.Classe) Player {
	return newPlayer(nom, 1, classe)
}

func NewLimule() Player {
	return newPlayer("Limule", 3, classes.NewMage()) // 130 PV, comme avant
}

// GainXP ajoute de l'XP et enchaîne les montées de niveau si nécessaire.
func (p *Player) GainXP(amount int) {
	if p.Level >= MaxLevel {
		return
	}
	p.Experience += amount
	fmt.Printf("+%d Experience\n", amount)

	for p.Level < MaxLevel && p.Experience >= XPToNextLevel(p.Level) {
		p.Experience -= XPToNextLevel(p.Level)
		p.levelUp()
	}
}

func (p *Player) levelUp() {
	p.Level++
	p.MaxHealth = MaxHealthForLevel(p.Level)
	p.Health = p.MaxHealth // soin complet à la montée de niveau

	fmt.Println("|--------------------------|")
	fmt.Printf(" NIVEAU SUPÉRIEUR ! Niveau %d\n", p.Level)
	fmt.Printf(" PV max : %d\n", p.MaxHealth)
	fmt.Println("|--------------------------|")
}

func (p Player) Afficher() {
	fmt.Println("|--------------------------|")
	fmt.Printf(" Nom      : %s\n", p.Name)
	fmt.Printf(" Classe   : %s\n", p.Class.Nom)
	fmt.Printf(" Niveau   : %d (Experience : %d/%d)\n", p.Level, p.Experience, XPToNextLevel(p.Level))
	fmt.Printf(" Vie      : %d/%d\n", p.Health, p.MaxHealth)
	fmt.Printf(" Argent   : %d\n", p.Money)
	fmt.Printf(" Mana     : %d\n", p.Class.Mana)
	fmt.Printf(" Attaque Physique : %d / Attaque Magique : %d\n", p.Class.AttaquePhysique, p.Class.AttaqueMagique)
	fmt.Println(" Compétences :")
	for _, c := range p.Class.Competences {
		fmt.Printf("   - %s (%d dégâts)\n", c.Nom, c.Degats)
	}
	fmt.Println("|--------------------------|")
}

// IsDead : receveur pointeur, sinon la réapparition ne modifie qu'une copie.
func (p *Player) IsDead() bool {
	if p.Health <= 0 {
		fmt.Println("Tu es mort, réapparition avec 50% de ta vie max")
		p.Health = p.MaxHealth / 2
		return true
	}
	return false
}
