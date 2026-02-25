package game

type EventType string

const (
	EventAttack     EventType = "attack"
	EventHeal       EventType = "heal"
	EventLevelUp    EventType = "levelup"
	EventLoot       EventType = "loot"
	EventDeath      EventType = "death"
	EventExperience EventType = "experience"
)

type Event struct {
	Type    EventType   `json:"type"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type AttackData struct {
	Attacker string `json:"attacker"`
	Target   string `json:"target"`
	Damage   int    `json:"damage"`
	Attack   string `json:"attack_name"`
}

type LootData struct {
	MonsterName string  `json:"monster_name"`
	Items       []*Slot `json:"items"`
}
