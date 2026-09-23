package potion

import (
	"fmt"

	player "piscine/player"
)

func TakePot(p *player.Player) bool {
	if p == nil {
		return false
	}

	seuil := p.MaxHealth * 20 / 100
	if p.Health > seuil {
		return false
	}

	indexPotion := -1
	for i, item := range p.Inventory.Items {
		if item.Nom == "potion de soin" {
			indexPotion = i
			break
		}
	}

	if indexPotion == -1 {
		fmt.Println("Tu as moins de 20% de vie, mais tu n'as pas de potion de soin.")
		return false
	}

	fmt.Println("Tu as moins de 20% de vie ! Utilisation d'une potion de soin.")

	p.Inventory.Items = append(
		p.Inventory.Items[:indexPotion],
		p.Inventory.Items[indexPotion+1:]...,
	)

	healAmount := p.MaxHealth / 2
	if p.Health+healAmount > p.MaxHealth {
		p.Health = p.MaxHealth
	} else {
		p.Health += healAmount
	}

	fmt.Printf("Tu as regagné %d PV (50%% de ta vie max). Vie actuelle : %d/%d\n", healAmount, p.Health, p.MaxHealth)
	return true
}

func UsePoisonPotion(p *player.Player, targetHP *int, maxTargetHP int) bool {
	if p == nil || targetHP == nil || maxTargetHP <= 0 {
		return false
	}

	indexPotion := -1
	for i, item := range p.Inventory.Items {
		if item.Nom == "potion de poison" {
			indexPotion = i
			break
		}
	}

	if indexPotion == -1 {
		fmt.Println("Tu n'as pas de potion de poison.")
		return false
	}

	p.Inventory.Items = append(
		p.Inventory.Items[:indexPotion],
		p.Inventory.Items[indexPotion+1:]...,
	)

	fmt.Println("Tu utilises une potion de poison !")

	damage := maxTargetHP / 100
	if damage < 1 {
		damage = 1
	}

	if *targetHP-damage < 0 {
		*targetHP = 0
	} else {
		*targetHP -= damage
	}

	fmt.Printf("L'ennemi perd %d PV (1%% de sa vie). Vie restante : %d\n", damage, *targetHP)

	return true
}

// UseManaPotion consomme une potion de mana et restaure 50% du mana maximum.
func UseManaPotion(p *Player) bool {
	if p == nil {
		return false
	}

	indexPotion := -1
	for i, item := range p.Inventory.Items {
		if item.Nom == "potion de mana" {
			indexPotion = i
			break
		}
	}

	if indexPotion == -1 {
		fmt.Println("Tu n'as pas de potion de mana.")
		return false
	}

	p.Inventory.Items = append(
		p.Inventory.Items[:indexPotion],
		p.Inventory.Items[indexPotion+1:]...,
	)

	recovery := p.MaxMana / 2
	p.Mana += recovery
	if p.Mana > p.MaxMana {
		p.Mana = p.MaxMana
	}

	fmt.Printf("Tu as regagné %d mana (50%% de ton mana max). Mana actuel : %d/%d\n", recovery, p.Mana, p.MaxMana)
	return true
}
