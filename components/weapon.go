package components

import "github.com/yohamta/donburi"

type WeaponComponent struct {
	Name          string
	MinimumDamage int
	MaximumDamage int
	ToHitBonus    int
}

var Weapon = donburi.NewComponentType[WeaponComponent]()
