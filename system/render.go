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
	query *donburi.Query
}

var Render = &render{
	query: ecs.NewQuery(
		layers.RealWorld,
		filter.Contains(
			components.Position,
			components.Renderable,
		)),
}

func (r *render) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	r.query.Each(ecs.World, func(entry *donburi.Entry) {
		position := components.Position.Get(entry)
		sprite := components.Renderable.Get(entry)

		op := &ebiten.DrawImageOptions{}
		sw, sh := float64(screen.Bounds().Dx()), float64(screen.Bounds().Dy())
		op.GeoM.Translate(float64(position.X)*sw, float64(position.Y)*sh)
		screen.DrawImage(sprite.Image, op)
	})
}
