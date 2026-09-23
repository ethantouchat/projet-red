package weapon

type TypeWeapon struct {
	Name        string
	AttackBonus int
	Type        string
	Price       int
	Rarity      string
	Niveau      int
}

func NewArmeDeClasse(classe string) *TypeWeapon {
	switch classe {
	case "Mage":
		return &TypeWeapon{Name: "Bâton de Chêne", AttackBonus: 3, Type: "magique", Price: 0, Rarity: "Commun", Niveau: 1}
	case "Epeiste":
		return &TypeWeapon{Name: "Épée Courte", AttackBonus: 4, Type: "physique", Price: 0, Rarity: "Commun", Niveau: 1}
	case "Assassin":
		return &TypeWeapon{Name: "Dague", AttackBonus: 3, Type: "physique", Price: 0, Rarity: "Commun", Niveau: 1}
	default:
		return &TypeWeapon{Name: "Poing", AttackBonus: 1, Type: "physique", Price: 0, Rarity: "Commun", Niveau: 1}
	}
}

func ListeArmes() []TypeWeapon {
	return []TypeWeapon{
		{Name: "Bâton de Chêne", AttackBonus: 3, Type: "magique", Price: 100, Rarity: "Commun", Niveau: 1},
		{Name: "Épée Courte", AttackBonus: 4, Type: "physique", Price: 120, Rarity: "Commun", Niveau: 1},
		{Name: "Dague Empoisonnée", AttackBonus: 5, Type: "physique", Price: 200, Rarity: "Rare", Niveau: 1},
		{Name: "Bâton de Fermon", AttackBonus: 6, Type: "magique", Price: 300, Rarity: "Rare", Niveau: 2},
		{Name: "Épée Longue", AttackBonus: 7, Type: "physique", Price: 400, Rarity: "Rare", Niveau: 2},
		{Name: "Marteau de Guerre", AttackBonus: 9, Type: "physique", Price: 500, Rarity: "Épique", Niveau: 3},
		{Name: "Bâton de L'Archimage", AttackBonus: 10, Type: "magique", Price: 600, Rarity: "Épique", Niveau: 3},
		{Name: "Épée Légende", AttackBonus: 12, Type: "physique", Price: 800, Rarity: "Légende", Niveau: 5},
		{Name: "Bâton de l'Implosion", AttackBonus: 14, Type: "magique", Price: 1000, Rarity: "Légende", Niveau: 5},
	}
}

func AmeliorerArme(w *TypeWeapon) bool {
	if w == nil {
		return false
	}

	w.AttackBonus += 2
	w.Niveau++

	return true
}
