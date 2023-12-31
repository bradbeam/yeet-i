package game

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"

	"github.com/bradbeam/yeet-i/components"
	"github.com/bradbeam/yeet-i/layers"
)

func (g *Game) EnterWorld() {
	levelQuery := ecs.NewQuery(
		layers.Floor,
		filter.Contains(
			components.Level,
		),
	)

	x, y := 0, 0
	levelEntity, ok := levelQuery.First(g.ecs.World)
	if ok {
		level := components.Level.Get(levelEntity)

		// Do we need to nil check here?
		x, y = level.Rooms[0].Center()
	}

	playerEntity := g.ecs.World.Entry(
		g.ecs.Create(
			layers.RealWorld,
			components.Player,
			components.Renderable,
			components.Movable,
			components.Position,
			components.Health,
			components.Weapon,
			components.Armor,
			components.Name,
			components.Message,
		),
	)

	// This seems like a bit of a chore :/
	donburi.SetValue(
		playerEntity,
		components.Renderable,
		components.RenderableComponent{
			Image: g.loadAssetFromFS("assets/player.png"),
		},
	)

	donburi.SetValue(
		playerEntity,
		components.Position,
		components.PositionComponent{
			X: x,
			Y: y,
		},
	)

	donburi.SetValue(
		playerEntity,
		components.Health,
		components.HealthComponent{
			MaxHealth:     30,
			CurrentHealth: 30,
		},
	)

	donburi.SetValue(
		playerEntity,
		components.Weapon,
		components.WeaponComponent{
			Name:          "Battle Axe",
			MinimumDamage: 10,
			MaximumDamage: 20,
			ToHitBonus:    3,
		},
	)

	donburi.SetValue(
		playerEntity,
		components.Armor,
		components.ArmorComponent{
			Name:       "Plate Armor",
			Defense:    15,
			ArmorClass: 18,
		},
	)

	donburi.SetValue(
		playerEntity,
		components.Name,
		components.NameComponent{
			Label: "Ohai",
		},
	)

	donburi.SetValue(
		playerEntity,
		components.Message,
		components.MessageComponent{
			AttackMessage:    "",
			DeadMessage:      "",
			GameStateMessage: "",
		},
	)

	// spawn enemies
}
