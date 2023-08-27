package scene

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type DungeonScene struct {
	ecs *ecs.ECS
}

func (ps *DungeonScene) Update() {
	ps.ecs.Update()
}

func (ps *DungeonScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{20, 20, 40, 255})
	ps.ecs.Draw(screen)
}

func NewDungeon() *DungeonScene {
	ecs := ecs.NewECS(donburi.NewWorld())

	return &DungeonScene{
		ecs: ecs,
	}
}
