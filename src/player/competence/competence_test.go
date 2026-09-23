package competence

import "testing"

func TestThunderboltOnlyLimuleCanUseIt(t *testing.T) {
	thunderbolt := Competence("Thunderbolt")

	if !CanUse(thunderbolt, "Limule", 8) {
		t.Fatal("Limule doit pouvoir utiliser Thunderbolt au niveau requis")
	}

	if CanUse(thunderbolt, "Player", 8) {
		t.Fatal("Thunderbolt ne doit pas être utilisable par un autre joueur")
	}
}

func TestOtherCompetencesAreAvailableToPlayers(t *testing.T) {
	fireball := Competence("Fireball")

	if !CanUse(fireball, "Player", 1) {
		t.Fatal("Fireball doit être utilisable par un joueur au niveau requis")
	}
}