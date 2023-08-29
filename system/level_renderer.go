package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"

	"github.com/bradbeam/yeet-i/components"
	"github.com/bradbeam/yeet-i/layers"
)

type levelRender struct {
	query *donburi.Query
}

var LevelRender = &levelRender{
	query: ecs.NewQuery(
		layers.Floor,
		filter.Contains(
			components.Level,
			components.Renderable,
		)),
}

func (r *levelRender) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	r.query.Each(ecs.World, func(entry *donburi.Entry) {
		level := components.Level.Get(entry)
		level.Draw(screen)
	})
}
