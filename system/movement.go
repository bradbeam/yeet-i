package system

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	// TPS is 60 by default, so we just need something to go
	// at least that high
	ticks uint8
	rate  uint8
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

	rate: 30,
}

func (m *movement) Update(ecs *ecs.ECS) {
	m.ticks++

	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		if m.rate < 255 {
			m.rate++
		}

		fmt.Println(m.rate)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		if m.rate > 0 {
			m.rate--
		}

		fmt.Println(m.rate)
	}

	if m.ticks%m.rate != 0 {
		return
	}

	// Reset counter so we can do arbitrary mod's above
	// and not worry about periodically skipping a move
	m.ticks = 0

	x := 0
	y := 0

	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		y = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		y = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		x = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
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
