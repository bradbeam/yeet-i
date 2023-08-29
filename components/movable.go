package components

import "github.com/yohamta/donburi"

type MovableComponent struct{}

var Movable = donburi.NewComponentType[MovableComponent]()
