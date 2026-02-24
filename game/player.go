package game

import (
	"fmt"
)

type Player struct {
	Name  string
	Class Class
	HP    int
	Exp   int
	Level int
}

type Class struct {
	Name      string
	Abilities []string
	Role      string
}

func (p *Player) ModifyExp(exp int) {
	p.Exp += exp
	fmt.Printf("%s has gained %d exp!\n", p.Name, exp)

	if p.Exp > 100 {
		p.Level += 1
		p.Exp = p.Exp - 100
		fmt.Printf("%s has leveled up to level %d\n", p.Name, p.Level)
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
	}

	fmt.Printf("%s attacks %s and deals %d damage\n", p.Name, m.Name, dmg)
	fmt.Printf("%s has %d hp remaining\n", m.Name, m.HP)
}
