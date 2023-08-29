package components

import (
	"github.com/yohamta/donburi"
)

type HealthComponent struct {
	MaxHealth     int
	CurrentHealth int
	// RegenRate?
	// HealthBonuses?
}

var Health = donburi.NewComponentType[HealthComponent]()
