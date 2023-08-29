package components

import "github.com/yohamta/donburi"

type ArmorComponent struct {
	Name       string
	Defense    int
	ArmorClass int
}

var Armor = donburi.NewComponentType[ArmorComponent]()
