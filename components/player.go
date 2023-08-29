package components

import "github.com/yohamta/donburi"

type PlayerComponent struct{}

var Player = donburi.NewComponentType[PlayerComponent]()
