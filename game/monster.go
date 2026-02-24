package game

import (
	"fmt"
	"math/rand/v2"
)

type Monster struct {
	Name    string
	HP      int
	Exp     int
	Level   int
	Attacks []string
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

		fmt.Printf("%s slashes %s and deals %d\n", m.Name, p.Name, dmg)
		fmt.Printf("%s has %d hp remaining\n", p.Name, p.HP)
	} else if currentAttack == "snatch" {
		dmg = rand.IntN(5)
		p.HP -= dmg

		if p.HP <= 0 {
			p.HP = 0
		}

		fmt.Printf("%s snatch %s's items and deals %d\n", m.Name, p.Name, dmg)
		fmt.Printf("%s has %d hp remaining\n", p.Name, p.HP)
	} else {
		dmg = rand.IntN(10)
		p.HP -= dmg

		if p.HP <= 0 {
			p.HP = 0
		}

		fmt.Printf("%s attacks %s and deals %d\n", m.Name, p.Name, dmg)
		fmt.Printf("%s has %d hp remaining\n", p.Name, p.HP)
	}
}
