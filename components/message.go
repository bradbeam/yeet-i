package components

import "github.com/yohamta/donburi"

type MessageComponent struct {
	AttackMessage    string
	DeadMessage      string
	GameStateMessage string
}

var Message = donburi.NewComponentType[MessageComponent]()
