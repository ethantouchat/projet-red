package competence

type Skill struct {
	Name          string
	Class         string
	Type          string
	ManaCost      int
	AttackBonus   int
	RequiredLevel int
	Description   string
	xp            int
	level         int
}

func NewSkill(name, classe, skillType string, manaCost, attackBonus, requiredLevel int, description string) Skill {
	return Skill{
		Name:          name,
		Class:         classe,
		Type:          skillType,
		ManaCost:      manaCost,
		AttackBonus:   attackBonus,
		RequiredLevel: requiredLevel,
		Description:   description,
	}
}

func (s *Skill) GainXP(amount int) bool {
	if s == nil || amount <= 0 {
		return false
	}

	s.xp += amount
	leveledUp := false
	for s.xp >= 50 {
		s.xp -= 50
		s.level++
		s.AttackBonus += 1
		leveledUp = true
	}
	return leveledUp
}

func (s *Skill) LevelUp() bool {
	if s == nil {
		return false
	}
	return s.GainXP(50)
}

func (s *Skill) Level() int {
	if s == nil {
		return 0
	}
	return s.level
}

func (s *Skill) XP() int {
	if s == nil {
		return 0
	}
	return s.xp
}

func Competence(name string) Skill {
	switch name {
	case "Fireball":
		return NewSkill("Fireball", "Mage", "magique", 10, 5, 1, "Lance une boule de feu sur l'ennemi.")
	case "Mana Detection":
		return NewSkill("Mana Detection", "Mage", "magique", 3, 0, 1, "Révèle le mana de l'ennemi.")
	case "Stealth":
		return NewSkill("Stealth", "Assassin", "physique", 5, 3, 1, "Rend l'assassin plus difficile à détecter.")
	case "Dagger Mastery":
		return NewSkill("Dagger Mastery", "Assassin", "physique", 0, 2, 1, "Améliore les dégâts infligés avec les dagues.")
	case "Vertical Strike":
		return NewSkill("Vertical Strike", "Swordsman", "physique", 2, 3, 1, "Effectue une frappe verticale puissante.")
	case "Horizontal Strike":
		return NewSkill("Horizontal Strike", "Swordsman", "physique", 3, 3, 1, "Effectue une frappe horizontale large.")
	case "Flame Emperor":
		return NewSkill("Flame Emperor", "Mage", "magique", 25, 35, 10, "Déchaîne une attaque de feu puissante.")
	case "Predator Aura":
		return NewSkill("Predator Aura", "Assassin", "physique", 0, 10, 5, "Augmente la puissance d'attaque de l'utilisateur.")
	case "Advanced Blade":
		return NewSkill("Advanced Blade", "Swordsman", "physique", 0, 15, 8, "Améliore les dégâts des attaques d'épée.")
	case "Thunderbolt":
		return NewSkill("Thunderbolt", "Mage", "magique", 20, 25, 8, "Frappe un ennemi de la foudre.")
	default:
		return Skill{}
	}
}

func CanUse(skill Skill, playerLevel int) bool {
	return skill.Name != "" && playerLevel >= skill.RequiredLevel
}

func Level(skill Skill, playerLevel int, playerLevelUp func(), competenceLevelUp func()) {
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
