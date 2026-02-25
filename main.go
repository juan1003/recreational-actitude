package main

import (
	"fmt"
	"math/rand/v2"
	"recreational-actitude/game"
	"recreational-actitude/network"
	"time"
)

func main() {
	server := network.NewServer()
	go server.Start(":8080")

	// Wait a bit for server to start
	time.Sleep(1 * time.Second)

	player1 := game.Player{
		Name: "Gabriel Branford",
		HP:   1000000,
		Class: game.Class{
			Name:      "Gladiator",
			Abilities: []string{"cover", "holy"},
			Role:      "tank",
		},
		Exp:       0,
		Level:     1,
		Inventory: game.NewInventory(game.Backpack, 10),
		OnEvent:   server.BroadcastEvent,
	}

	boss1 := game.Monster{
		Name:    "Super Dimension Fortress Macross",
		HP:      15000,
		Exp:     rand.IntN(2500),
		Level:   10,
		Attacks: []string{"attack", "Starcross", "Hydro Cannons", "Frenzy Shot"},
		LootTable: []game.LootEntry{
			{
				Item: &game.Item{
					ID:          "1",
					Name:        "Potion",
					Description: "Restores HP",
					Type:        game.Consumable,
					MaxStack:    10,
				},
				Chance:      0.8,
				MinQuantity: 1,
				MaxQuantity: 3,
			},
			{
				Item: &game.Item{
					ID:          "2",
					Name:        "Steel Plate",
					Description: "A piece of armor material",
					Type:        game.Material,
					MaxStack:    1,
				},
				Chance:      0.2,
				MinQuantity: 1,
				MaxQuantity: 1,
			},
		},
		OnEvent: server.BroadcastEvent,
	}

	fmt.Println("Waiting for a client to connect before starting simulation...")
	// In a real scenario, you might wait for a "start" message from WS
	// For this prototype, let's just wait 5 seconds or until a manual trigger
	time.Sleep(5 * time.Second)

	for i := 1; i < 1000; i++ {
		boss1.AttackPlayer(&player1)
		time.Sleep(500 * time.Millisecond)

		if player1.HP <= 0 {
			fmt.Printf("%s has defeated %s\n", boss1.Name, player1.Name)
			break
		}

		// Random damage for simulation
		dmg := rand.IntN(40)
		player1.AttackMonster(&boss1, dmg)
		time.Sleep(500 * time.Millisecond)

		if boss1.HP <= 0 {
			// AttackMonster already handles exp and loot
			fmt.Printf("Player Status: %s | %s Level: %d | Exp: %d\n", player1.Name, player1.Class.Name, player1.Level, player1.Exp)
			player1.Inventory.Print()
			break
		}
	}

	// Keep the server running after simulation ends
	select {}
}
