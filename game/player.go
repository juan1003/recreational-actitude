package game

import (
	"fmt"
)

type Player struct {
	Name      string
	Class     Class
	HP        int
	Exp       int
	Level     int
	Inventory *Inventory
	OnEvent   func(Event)
}

func (p *Player) emit(event Event) {
	if p.OnEvent != nil {
		p.OnEvent(event)
	}
}

type Class struct {
	Name      string
	Abilities []string
	Role      string
}

func (p *Player) ModifyExp(exp int) {
	p.Exp += exp
	msg := fmt.Sprintf("%s has gained %d exp!", p.Name, exp)
	p.emit(Event{Type: EventExperience, Message: msg, Data: exp})
	fmt.Println(msg)

	if p.Exp > 100 {
		p.Level += 1
		p.Exp = p.Exp - 100
		levelMsg := fmt.Sprintf("%s has leveled up to level %d", p.Name, p.Level)
		p.emit(Event{Type: EventLevelUp, Message: levelMsg, Data: p.Level})
		fmt.Println(levelMsg)
	}
}

func (c *Class) HasAbility(search string) bool {
	for _, a := range c.Abilities {
		if a == search {
			return true
		}
	}
	return false
}

func (p *Player) AttackMonster(m *Monster, dmg int) {
	m.HP -= dmg

	if m.HP <= 0 {
		m.HP = 0
		deathMsg := fmt.Sprintf("%s attacks %s and deals %d damage. %s has been defeated!", p.Name, m.Name, dmg, m.Name)
		p.emit(Event{
			Type:    EventAttack,
			Message: deathMsg,
			Data: AttackData{
				Attacker: p.Name,
				Target:   m.Name,
				Damage:   dmg,
				Attack:   "attack",
			},
		})
		fmt.Println(deathMsg)
		p.ModifyExp(m.Exp)

		drops := m.DropLoot()
		if len(drops) > 0 {
			lootMsg := fmt.Sprintf("%s dropped loot!", m.Name)
			p.emit(Event{
				Type:    EventLoot,
				Message: lootMsg,
				Data: LootData{
					MonsterName: m.Name,
					Items:       drops,
				},
			})
			fmt.Printf("%s dropped:\n", m.Name)
			for _, slot := range drops {
				fmt.Printf("- %s x%d\n", slot.Item.Name, slot.Quantity)
				if p.Inventory != nil {
					err := p.Inventory.AddItem(slot.Item, slot.Quantity)
					if err != nil {
						fmt.Printf("Could not add item to inventory: %v\n", err)
					}
				}
			}
		}
	} else {
		attackMsg := fmt.Sprintf("%s attacks %s and deals %d damage", p.Name, m.Name, dmg)
		p.emit(Event{
			Type:    EventAttack,
			Message: attackMsg,
			Data: AttackData{
				Attacker: p.Name,
				Target:   m.Name,
				Damage:   dmg,
				Attack:   "attack",
			},
		})
		fmt.Println(attackMsg)
		fmt.Printf("%s has %d hp remaining\n", m.Name, m.HP)
	}
}
