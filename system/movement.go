package system

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"

	"github.com/bradbeam/yeet-i/components"
)

type movement struct {
	query *donburi.Query
}

var Movement = &movement{
	query: donburi.NewQuery(
		filter.Contains(
			components.Player,
			components.Position,
		)),
}

func (m *movement) Update(ecs *ecs.ECS) {
	x := 0
	y := 0
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		y = -1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		y = 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		x = -1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		x = 1
	}

	if x == 0 && y == 0 {
		return
	}

	m.query.Each(ecs.World, func(entry *donburi.Entry) {
		position := components.Position.Get(entry)

		position.X += x
		position.Y += y

		fmt.Printf("%+v\n", position)
	})
}