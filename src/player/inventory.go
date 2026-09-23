package player

import (
	"errors"
	"fmt"
)

const DefaultCapacity = 10

var (
	ErrInventoryFull   = errors.New("inventaire plein")
	ErrInvalidSlot     = errors.New("emplacement invalide")
	ErrInvalidCapacity = errors.New("la capacité doit augmenter d'une valeur positive")
)

type Item struct {
	Nom string
}

type Inventory struct {
	Items    []Item
	Capacity int
}

func NewInventory() Inventory {
	return Inventory{Items: []Item{}, Capacity: DefaultCapacity}
}

func (inv *Inventory) IsFull() bool {
	return len(inv.Items) >= inv.Capacity
}

func (inv *Inventory) AddItem(nom string) error {
	if inv.IsFull() {
		return ErrInventoryFull
	}
	inv.Items = append(inv.Items, Item{Nom: nom})
	return nil
}
func (inv *Inventory) SwapItem(index int, nom string) (Item, error) {
	if index < 0 || index >= len(inv.Items) {
		return Item{}, ErrInvalidSlot
	}
	old := inv.Items[index]
	inv.Items[index] = Item{Nom: nom}
	return old, nil
}

func (inv *Inventory) IncreaseCapacity(n int) error {
	if n <= 0 {
		return ErrInvalidCapacity
	}
	inv.Capacity += n
	return nil
}

func (inv Inventory) Afficher() {
	fmt.Printf("Inventaire (%d/%d)\n", len(inv.Items), inv.Capacity)
	if len(inv.Items) == 0 {
		fmt.Println("Inventaire vide")
		return
	}
	if inv.IsFull() {
		fmt.Println("Inventaire plein")
	}
	for i, item := range inv.Items {
		fmt.Printf(" %d - %s\n", i, item.Nom)
	}
}

func (inv Inventory) CountItem(nom string) int {
	count := 0
	for _, item := range inv.Items {
		if item.Nom == nom {
			count++
		}
	}
	return count
}

func (inv *Inventory) RemoveItem(nom string) bool {
	for i, item := range inv.Items {
		if item.Nom == nom {
			inv.Items = append(inv.Items[:i], inv.Items[i+1:]...)
			return true
		}
	}
	return false
}

func (inv *Inventory) RemoveItems(nom string, count int) bool {
	removed := 0
	for i := 0; i < len(inv.Items); {
		if inv.Items[i].Nom == nom {
			inv.Items = append(inv.Items[:i], inv.Items[i+1:]...)
			removed++
			if removed >= count {
				break
			}
		} else {
			i++
		}
	}
	return removed >= count
}

