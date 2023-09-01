package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"

	"github.com/bradbeam/yeet-i/components"
	"github.com/bradbeam/yeet-i/maps"
)

// TODO Probably should split this between player and monster
type movement struct {
	query      *donburi.Query
	levelQuery *donburi.Query
}

var Movement = &movement{
	query: donburi.NewQuery(
		filter.Contains(
			components.Player,
			components.Position,
		)),

	levelQuery: donburi.NewQuery(
		filter.Contains(
			components.Level,
		),
	),
}

func (m *movement) Update(ecs *ecs.ECS) {
	x := 0
	y := 0

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		y = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		y = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		x = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		x = 1
	}

	if x == 0 && y == 0 {
		return
	}

	levelEntity, ok := m.levelQuery.First(ecs.World)
	if !ok {
		return
	}

	level := components.Level.Get(levelEntity)

	m.query.Each(ecs.World, func(entry *donburi.Entry) {
		position := components.Position.Get(entry)

		startingIndex := level.GetIndexFromXY(position.X, position.Y)

		destIndex := level.GetIndexFromXY(position.X+x, position.Y+y)
		if level.Tiles[destIndex].Tile.Blocked() {
			return
		}

		position.X += x
		position.Y += y

		level.Tiles[startingIndex].Tile.State = maps.TileStateFree
		level.Tiles[destIndex].Tile.State = maps.TileStateOccupied
	})
}
