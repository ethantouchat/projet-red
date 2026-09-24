package competence

import "fmt"

type Skill struct {
	Name          string
	Class         string
	Type          string
	ManaCost      int
	AttackBonus   int
	RequiredLevel int
	AllowedPlayer string
	Description   string
	xp            int
	level         int
}

func (s Skill) Level() int {
	return s.level
}

func (s *Skill) GainXP(amount int) bool {
	if s == nil || amount <= 0 {
		return false
	}

	s.xp += amount
	for s.xp >= 50 {
		s.xp -= 50
		s.level++
		s.AttackBonus += 1
		return true
	}
	return false
}

func (s *Skill) LevelUp() bool {
	if s == nil {
		return false
	}
	return s.GainXP(50)
}

func Competence(name string) Skill {
	switch name {
	case "Fireball":
		return Skill{
			Name:          "Fireball",
			Class:         "Mage",
			ManaCost:      10,
			RequiredLevel: 1,
			AllowedPlayer: "Player",
			Description:   "Launches a fireball at an enemy.",
			xp:            0,
			level:         0,
		}
	case "Mana Detection":
		return Skill{
			Name:          "Mana Detection",
			Class:         "Mage",
			RequiredLevel: 1,
			AllowedPlayer: "Player",
			Description:   "Reveals the mana of an enemy.",
			xp:            0,
			level:         0,
		}
	case "Stealth":
		return Skill{
			Name:          "Stealth",
			Class:         "Assassin",
			ManaCost:      5,
			RequiredLevel: 1,
			AllowedPlayer: "Player",
			Description:   "Makes the assassin harder to detect.",
			xp:            0,
			level:         0,
		}
	case "Dagger Mastery":
		return Skill{
			Name:          "Dagger Mastery",
			Class:         "Assassin",
			RequiredLevel: 1,
			AllowedPlayer: "Player",
			Description:   "Improves damage dealt with daggers.",
			xp:            0,
			level:         0,
		}
	case "Vertical Strike":
		return Skill{
			Name:          "Vertical Strike",
			Class:         "Swordsman",
			ManaCost:      2,
			RequiredLevel: 1,
			AllowedPlayer: "Player",
			Description:   "Performs a powerful vertical strike.",
			xp:            0,
			level:         0,
		}
	case "Horizontal Strike":
		return Skill{
			Name:          "Horizontal Strike",
			Class:         "Swordsman",
			ManaCost:      3,
			RequiredLevel: 1,
			AllowedPlayer: "Player",
			Description:   "Performs a wide horizontal strike.",
			xp:            0,
			level:         0,
		}
	case "Flame Emperor":
		return Skill{
			Name:          "Flame Emperor",
			Class:         "Mage",
			ManaCost:      25,
			AttackBonus:   35,
			RequiredLevel: 8,
			AllowedPlayer: "Player",
			Description:   "Unleashes a powerful flame attack.",
			xp:            0,
			level:         0,
		}
	case "Thunderbolt":
		return Skill{
			Name:          "Thunderbolt",
			Class:         "Mage",
			ManaCost:      20,
			AttackBonus:   25,
			RequiredLevel: 8,
			AllowedPlayer: "Limule",
			Description:   "Strikes an enemy with lightning.",
			xp:            0,
			level:         0,
		}
	case "Flash":
		return Skill{
			Name:          "Flash",
			Class:         "Mage",
			ManaCost:      10,
			AttackBonus:   10,
			RequiredLevel: 8,
			AllowedPlayer: "Limule",
			Description:   "Strikes an enemy big flash.",
			xp:            0,
			level:         0,
		}	
	case "Predator Aura":
		return Skill{
			Name:          "Predator Aura",
			Class:         "Assassin",
			AttackBonus:   10,
			RequiredLevel: 8,
			AllowedPlayer: "Player",
			Description:   "Increases the user's attack power.",
			xp:            0,
			level:         0,
		}
	case "Advanced Blade":
		return Skill{
			Name:          "Advanced Blade",
			Class:         "Swordsman",
			AttackBonus:   15,
			RequiredLevel: 8,
			AllowedPlayer: "Player",
			Description:   "Improves the damage of sword attacks.",
			xp:            0,
			level:         0,
		}
	default:
		return Skill{}
	}
}

func CanUse(skill Skill, playerName string, playerLevel int) bool {
	return skill.Name != "" &&
		(skill.AllowedPlayer == "" || skill.AllowedPlayer == playerName) &&
		playerLevel >= skill.RequiredLevel
}

func Level(skill Skill, playerName string, playerLevel int, playerLevelUp func(), competenceLevelUp func()) {
	if !CanUse(skill, playerName, playerLevel) {
		return
	}

	if playerLevel >= skill.RequiredLevel {
		if playerLevelUp != nil {
			playerLevelUp()
		}

		if competenceLevelUp != nil {
			competenceLevelUp()
		}
		return
	}

	if competenceLevelUp != nil {
		competenceLevelUp()
	}
}

func CanActivate(skill Skill, playerName string, playerLevel int, mana int) bool {
	if !CanUse(skill, playerName, playerLevel) {
		return false
	}

	if mana < skill.ManaCost {
		fmt.Printf("Mana insuffisant pour utiliser %s : %d mana nécessaire(s).\n", skill.Name, skill.ManaCost)
		return false
	}

	return true
}
