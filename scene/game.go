package scene

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/bradbeam/yeet-i/maps"
)

type GameScene struct {
	Level *maps.Level
}

func (g *GameScene) Update() {}

func (g *GameScene) Draw(screen *ebiten.Image) {
	g.Level.Draw(screen)
}
