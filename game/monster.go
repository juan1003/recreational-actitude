package game

import (
	"fmt"
	"math/rand/v2"
)

type LootEntry struct {
	Item        *Item
	Chance      float64 // 0.0 to 1.0
	MinQuantity int
	MaxQuantity int
}

type Monster struct {
	Name      string
	HP        int
	Exp       int
	Level     int
	Attacks   []string
	LootTable []LootEntry
	OnEvent   func(Event)
}

func (m *Monster) emit(event Event) {
	if m.OnEvent != nil {
		m.OnEvent(event)
	}
}

func (m *Monster) DropLoot() []*Slot {
	var drops []*Slot

	for _, entry := range m.LootTable {
		if rand.Float64() <= entry.Chance {
			quantity := entry.MinQuantity
			if entry.MaxQuantity > entry.MinQuantity {
				quantity += rand.IntN(entry.MaxQuantity - entry.MinQuantity + 1)
			}
			if quantity > 0 {
				drops = append(drops, &Slot{Item: entry.Item, Quantity: quantity})
			}
		}
	}

	return drops
}

func (m *Monster) AttackPlayer(p *Player) {

	var dmg int
	randomIndex := rand.IntN(len(m.Attacks))
	currentAttack := m.Attacks[randomIndex]

	if currentAttack == "slash" {
		dmg = rand.IntN(20)
		p.HP -= dmg

		if p.HP <= 0 {
			p.HP = 0
		}

		msg := fmt.Sprintf("%s slashes %s and deals %d", m.Name, p.Name, dmg)
		m.emit(Event{
			Type:    EventAttack,
			Message: msg,
			Data: AttackData{
				Attacker: m.Name,
				Target:   p.Name,
				Damage:   dmg,
				Attack:   currentAttack,
			},
		})
		fmt.Println(msg)
		fmt.Printf("%s has %d hp remaining\n", p.Name, p.HP)
	} else if currentAttack == "snatch" {
		dmg = rand.IntN(5)
		p.HP -= dmg

		if p.HP <= 0 {
			p.HP = 0
		}

		msg := fmt.Sprintf("%s snatch %s's items and deals %d", m.Name, p.Name, dmg)
		m.emit(Event{
			Type:    EventAttack,
			Message: msg,
			Data: AttackData{
				Attacker: m.Name,
				Target:   p.Name,
				Damage:   dmg,
				Attack:   currentAttack,
			},
		})
		fmt.Println(msg)
		fmt.Printf("%s has %d hp remaining\n", p.Name, p.HP)
	} else {
		dmg = rand.IntN(10)
		p.HP -= dmg

		if p.HP <= 0 {
			p.HP = 0
		}

		msg := fmt.Sprintf("%s attacks %s and deals %d", m.Name, p.Name, dmg)
		m.emit(Event{
			Type:    EventAttack,
			Message: msg,
			Data: AttackData{
				Attacker: m.Name,
				Target:   p.Name,
				Damage:   dmg,
				Attack:   "attack",
			},
		})
		fmt.Println(msg)
		fmt.Printf("%s has %d hp remaining\n", p.Name, p.HP)
	}
}
