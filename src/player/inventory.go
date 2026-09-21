package player

import "fmt"

type Item struct {
	Nom string
}

type Inventory struct {
	Items []Item
}

func NewInventory() Inventory {
	return Inventory{Items: []Item{}}
}

func (inv *Inventory) AddItem(nom string) {
	inv.Items = append(inv.Items, Item{Nom: nom})
}

func (inv Inventory) Afficher() {
	if len(inv.Items) == 0 {
		fmt.Println("Inventaire vide")
		return
	} else if len(inv.Items) == 10 {
		fmt.Println("Inventaire plein")
	}
	for _, item := range inv.Items {
		fmt.Println(" -", item.Nom)
	}
}
