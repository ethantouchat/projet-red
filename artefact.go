package piscine 

type Equipment struct {
    Name   string
    Rank   string
    Life   int
    Attack int
    Mana   int
}
 
 func NewEquipment(name string) Equipment {
 	switch name {
 	case "Orb of Avarice":
 		return Equipment{Name: "Orb of Avarice", Rank: "SS", Attack: 2}
 	case "Goblin Ring":
 		return Equipment{Name: "Goblin Ring", Rank: "B", Mana: 17}
 	case "Traditional Goblin Necklace":
 		return Equipment{Name: "Traditional Goblin Necklace", Rank: "C", Life: 20, Attack: 2}
 	case "Goblin Helmet":
 		return Equipment{Name: "Goblin Helmet", Rank: "C", Life: 30, Attack: 5}
 	default:
 		return Equipment{}
 	}
 }
 
	 func ApplyEquipmentBonuses(attack, health, mana int, helmetEquipped, necklaceEquipped, ringEquipped, orbEquipped bool) (int, int, int) {
		 if helmetEquipped && necklaceEquipped {
			 attack = attack * 110 / 100
			 health = health * 110 / 100
		 }
 
		 if ringEquipped && orbEquipped {
			 mana = mana * 90 / 100
			 health = health * 90 / 100
		 }
 
		 return attack, health, mana
	 }
	}
}


















}