package components

import "github.com/yohamta/donburi"

type MonsterComponent struct{}

var Monster = donburi.NewComponentType[MonsterComponent]()
