package game

import (
	"errors"
	"fmt"
)

type ItemType string

const (
	Consumable ItemType = "consumable"
	Equipment  ItemType = "equipment"
	Material   ItemType = "material"
)

type Item struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Type        ItemType `json:"type"`
	MaxStack    int      `json:"max_stack"`
}

type Slot struct {
	Item     *Item `json:"item"`
	Quantity int   `json:"quantity"`
}

type InventoryType string

const (
	Backpack InventoryType = "backpack"
	Storage  InventoryType = "storage"
)

type Inventory struct {
	Type     InventoryType `json:"type"`
	Slots    []Slot        `json:"slots"`
	Capacity int           `json:"capacity"`
}

func NewInventory(invType InventoryType, capacity int) *Inventory {
	return &Inventory{
		Type:     invType,
		Slots:    make([]Slot, capacity),
		Capacity: capacity,
	}
}

func (inv *Inventory) AddItem(item *Item, amount int) error {
	remaining := amount

	// First pass: try to stack with existing items
	if item.MaxStack > 1 {
		for i := 0; i < len(inv.Slots); i++ {
			if inv.Slots[i].Item != nil && inv.Slots[i].Item.ID == item.ID {
				spaceInSlot := item.MaxStack - inv.Slots[i].Quantity
				if spaceInSlot > 0 {
					toAdd := spaceInSlot
					if remaining < toAdd {
						toAdd = remaining
					}
					inv.Slots[i].Quantity += toAdd
					remaining -= toAdd
				}
			}
			if remaining <= 0 {
				return nil
			}
		}
	}

	// Second pass: find empty slots
	for i := 0; i < len(inv.Slots); i++ {
		if inv.Slots[i].Item == nil {
			toAdd := item.MaxStack
			if remaining < toAdd {
				toAdd = remaining
			}
			inv.Slots[i].Item = item
			inv.Slots[i].Quantity = toAdd
			remaining -= toAdd
		}
		if remaining <= 0 {
			return nil
		}
	}

	if remaining < amount {
		return fmt.Errorf("inventory partially full: could not add %d of %s", remaining, item.Name)
	}

	return errors.New("inventory full")
}

func (inv *Inventory) Print() {
	fmt.Printf("--- %s Inventory (%d/%d) ---\n", inv.Type, inv.occupiedSlots(), inv.Capacity)
	for i, slot := range inv.Slots {
		if slot.Item != nil {
			fmt.Printf("[%d] %s x%d (%s)\n", i, slot.Item.Name, slot.Quantity, slot.Item.Type)
		} else {
			fmt.Printf("[%d] Empty\n", i)
		}
	}
}

func (inv *Inventory) occupiedSlots() int {
	count := 0
	for _, slot := range inv.Slots {
		if slot.Item != nil {
			count++
		}
	}
	return count
}
