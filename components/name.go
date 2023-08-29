package components

import "github.com/yohamta/donburi"

type NameComponent struct {
	Label string
}

var Name = donburi.NewComponentType[NameComponent]()
