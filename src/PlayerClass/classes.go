package classes

type Competence struct {
	Nom    string
	Degats int
}

type Classe struct {
	Nom             string
	Mana            int
	AttaquePhysique int
	AttaqueMagique  int
	Competences     []Competence
}

func NewMage() Classe {
	return Classe{
		Nom:             "Mage",
		Mana:            10,
		AttaquePhysique: 1,
		AttaqueMagique:  10,
		Competences: []Competence{
			{Nom: "Boule de feu", Degats: 3},
			{Nom: "Détection de Mana", Degats: 0},
		},
	}
}

func NewEpeiste() Classe {
	return Classe{
		Nom:             "Epeiste",
		Mana:            1,
		AttaquePhysique: 10,
		AttaqueMagique:  1,
		Competences: []Competence{
			{Nom: "Frappe Horizontale", Degats: 2},
			{Nom: "Frappe Verticale", Degats: 2},
		},
	}
}

func NewAssassin() Classe {
	return Classe{
		Nom:             "Assassin",
		Mana:            5,
		AttaquePhysique: 4,
		AttaqueMagique:  3,
		Competences: []Competence{
			{Nom: "Mille Entailles", Degats: 4},
			{Nom: "Furtivité", Degats: 0},
		},
	}
}

// Le 2e résultat false si la classe n'existe pas.
func GetClasse(nom string) (Classe, bool) {
	switch nom {
	case "mage":
		return NewMage(), true
	case "epeiste":
		return NewEpeiste(), true
	case "assassin":
		return NewAssassin(), true
	}
	return Classe{}, false
}
