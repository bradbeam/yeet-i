package components

import (
	"github.com/yohamta/donburi"

	"github.com/bradbeam/yeet-i/maps"
)

type LevelComponent struct {
	*maps.Level
}

var Level = donburi.NewComponentType[LevelComponent]()
