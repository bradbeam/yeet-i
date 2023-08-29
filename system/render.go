package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"

	"github.com/bradbeam/yeet-i/components"
	"github.com/bradbeam/yeet-i/layers"
)

type render struct {
	query      *donburi.Query
	levelQuery *donburi.Query
}

var Render = &render{
	query: ecs.NewQuery(
		layers.RealWorld,
		filter.Contains(
			components.Position,
			components.Renderable,
		),
	),
	levelQuery: ecs.NewQuery(
		layers.Floor,
		filter.Contains(
			components.Level,
		),
	),
}

func (r *render) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	level, ok := r.levelQuery.First(ecs.World)
	if !ok {
		return
	}

	l := components.Level.Get(level)

	r.query.Each(ecs.World, func(entry *donburi.Entry) {
		position := components.Position.Get(entry)
		sprite := components.Renderable.Get(entry)

		idx := l.GetIndexFromXY(position.X, position.Y)
		tile := l.Tiles[idx]

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(tile.X), float64(tile.Y))
		screen.DrawImage(sprite.Image, op)
	})
}
