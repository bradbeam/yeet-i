package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/ecs"

	"github.com/bradbeam/yeet-i/layers"
)

type GameScene struct {
	ECS *ecs.ECS
}

func (g *GameScene) Update() {}

func (g *GameScene) Draw(screen *ebiten.Image) {
	screen.Clear()

	//for _, layer := range []ecs.LayerID{layers.Default, layers.Floor, layers.Wall, layers.RealWorld} {
	for _, layer := range []ecs.LayerID{layers.Floor, layers.RealWorld} {
		g.ECS.DrawLayer(layer, screen)
	}
}
