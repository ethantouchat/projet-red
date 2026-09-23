package artefact

type TypeArtefact struct {
	Name         string
	HealthBonus  int
	ManaBonus    int
	AttackBonus  int
	Price        int
	Rarity       string
	Niveau       int
}

func NouvelArtefactDeClasse(classe string) *TypeArtefact {
	switch classe {
	case "Mage":
		return &TypeArtefact{Name: "Amulette de Mana", ManaBonus: 3, AttackBonus: 1, Price: 0, Rarity: "Commun", Niveau: 1}
	case "Epeiste":
		return &TypeArtefact{Name: "Ceinture de Fer", HealthBonus: 15, AttackBonus: 1, Price: 0, Rarity: "Commun", Niveau: 1}
	case "Assassin":
		return &TypeArtefact{Name: "Anneau d'Agilité", HealthBonus: 5, AttackBonus: 2, Price: 0, Rarity: "Commun", Niveau: 1}
	default:
		return &TypeArtefact{Name: "Pendentif", HealthBonus: 5, Price: 0, Rarity: "Commun", Niveau: 1}
	}
}

func ListeArtefacts() []TypeArtefact {
	return []TypeArtefact{
		{Name: "Amulette de Mana", ManaBonus: 3, AttackBonus: 1, Price: 150, Rarity: "Commun", Niveau: 1},
		{Name: "Ceinture de Fer", HealthBonus: 15, AttackBonus: 1, Price: 120, Rarity: "Commun", Niveau: 1},
		{Name: "Anneau d'Agilité", HealthBonus: 5, AttackBonus: 2, Price: 200, Rarity: "Rare", Niveau: 1},
		{Name: "Médaille de Gardien", HealthBonus: 30, AttackBonus: 2, Price: 400, Rarity: "Rare", Niveau: 2},
		{Name: "Bâton de Sagesse", ManaBonus: 5, AttackBonus: 2, Price: 500, Rarity: "Épique", Niveau: 3},
		{Name: "Cotte de Mailles", HealthBonus: 50, AttackBonus: 3, Price: 800, Rarity: "Épique", Niveau: 4},
		{Name: "Couronne du Roi", HealthBonus: 80, ManaBonus: 10, AttackBonus: 5, Price: 1200, Rarity: "Légende", Niveau: 6},
	}
}
