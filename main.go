package main

import (
	"fmt" 
	"math/rand/v2"
)

type Player struct {
  name string
  class Class 
  hp int 
  exp int
  level int
}

type Class struct {
  name string
  abilities []string
  role string
}

type Monster struct {
  name string
  hp int 
  exp int 
  level int
  attacks []string
}

func (p *Player) ModifyExp(exp int) {
  defer fmt.Println(fmt.Sprintf("%s has gained %d exp!", p.name, exp))
  p.exp += exp

  if p.exp > 100 {
	  p.level += 1
	  p.exp = p.exp - 100
      fmt.Println(fmt.Sprintf("%s has leveled up to level %d", p.name, p.level))
  } 
}

func (c *Class) HasAbility(search string) bool {
	for _, a := range c.abilities {
	  if a == search {
		return true
	  }
	}
	return false
}

func(p *Player) AttackMonster (m *Monster) {
	dmg := 40
	
	if p.class.HasAbility("holy") {
		bonus := 20
		dmg += bonus
		m.hp -= dmg	
			
	    if m.hp <= 0 {
		   m.hp = 0
	    } 
	}
    
	defer fmt.Println(fmt.Sprintf("%s has %d hp remaining", m.name, m.hp))	
	defer fmt.Println(fmt.Sprintf("%s attacks %s and deals %d damage", p.name, m.name, dmg))
}

func (m *Monster) AttackPlayer(p *Player) {
   
   var dmg int

   randomIndex := rand.IntN(len(m.attacks))

   currentAttack := m.attacks[randomIndex]


   if currentAttack == "slash" {
	 dmg = 20
	 p.hp -= dmg   
     
	 if p.hp <= 0 {
		p.hp = 0
	 } 
	 
	 fmt.Println(fmt.Sprintf("%s slashes %s and deals %d", m.name, p.name, dmg))      
     fmt.Println(fmt.Sprintf("%s has %d hp remaining", p.name, p.hp))
   } else if currentAttack == "snatch" {
	 dmg = 0
	 p.hp -= dmg
	 
	 if p.hp <= 0 {
		p.hp = 0
	 } 
     
	 fmt.Println(fmt.Sprintf("%s snatch %s's items and deals %d", m.name, p.name, dmg))      
     fmt.Println(fmt.Sprintf("%s has %d hp remaining", p.name, p.hp))
   } else {
	 dmg = 10
	 p.hp -= dmg
	 
	 if p.hp <= 0 {
		p.hp = 0
	 } 
     
	 fmt.Println(fmt.Sprintf("%s attacks %s and deals %d", m.name, p.name, dmg))      
     fmt.Println(fmt.Sprintf("%s has %d hp remaining", p.name, p.hp))
   } 
 }

func main() {
	player1 := Player{ name: "Gabriel Branford", hp: 200, class: Class{ name: "Gladiator", abilities: []string{"cover", "holy"}, role: "tank"}, exp: 0, level: 1 }      
	goblin1 := Monster{ name: "Goblin", hp: 200, exp: rand.IntN(140), level: 1, attacks: []string{"attack", "snatch", "slash"} }	
    
  for i := 1; i < 10; i++ {
  	
 	goblin1.AttackPlayer(&player1)

	if goblin1.hp <= 0 {
		fmt.Println(fmt.Sprintf("%s has defeated %s", player1.name, goblin1.name))
  		player1.ModifyExp(goblin1.exp)
  		fmt.Println(fmt.Sprintf("Player: %s | %s Level: %d | Exp: %d", player1.name, player1.class.name, player1.level, player1.exp))
		break	
	}
  
	player1.AttackMonster(&goblin1)
	
	if player1.hp <= 0 {
		fmt.Println(fmt.Sprintf("%s has defeated %s", goblin1.name, player1.name))	
		break	
	}
}
 
}
