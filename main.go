package main

import (
	"fmt"
	"math/rand/v2"
	"recreational-actitude/game"
)

func main() {
	player1 := game.Player{
		Name: "Gabriel Branford",
		HP:   200,
		Class: game.Class{
			Name:      "Gladiator",
			Abilities: []string{"cover", "holy"},
			Role:      "tank",
		},
		Exp:   0,
		Level: 1,
	}

	goblin1 := game.Monster{
		Name:    "Goblin",
		HP:      200,
		Exp:     rand.IntN(140),
		Level:   1,
		Attacks: []string{"attack", "snatch", "slash"},
	}

	for i := 1; i < 100; i++ {
		goblin1.AttackPlayer(&player1)

		if goblin1.HP <= 0 {
			fmt.Printf("%s has defeated %s\n", player1.Name, goblin1.Name)
			player1.ModifyExp(goblin1.Exp)
			fmt.Printf("Player: %s | %s Level: %d | Exp: %d\n", player1.Name, player1.Class.Name, player1.Level, player1.Exp)
			break
		}

		if player1.HP <= 0 {
			fmt.Printf("%s has defeated %s\n", goblin1.Name, player1.Name)
			break
		}

		// Random damage for simulation
		dmg := rand.IntN(40)
		player1.AttackMonster(&goblin1, dmg)

		if goblin1.HP <= 0 {
			fmt.Printf("%s has defeated %s\n", player1.Name, goblin1.Name)
			player1.ModifyExp(goblin1.Exp)
			fmt.Printf("Player: %s | %s Level: %d | Exp: %d\n", player1.Name, player1.Class.Name, player1.Level, player1.Exp)
			break
		}
	}
}
