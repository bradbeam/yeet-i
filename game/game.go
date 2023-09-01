package game

import (
	"image"
	"io"
	"io/fs"
	"log"
	"os"

	_ "image/png"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomono"

	"github.com/bradbeam/yeet-i/components"
	"github.com/bradbeam/yeet-i/layers"
	"github.com/bradbeam/yeet-i/maps"
	"github.com/bradbeam/yeet-i/scene"
	"github.com/bradbeam/yeet-i/system"
)

type Game struct {
	activeScene scene.Scene
	font        font.Face

	titleScene scene.Scene
	gameScene  scene.Scene

	fs fs.FS

	ecs *ecs.ECS
}

func NewGame(fs fs.FS) *Game {
	// TODO this needs a fair amount of cleanup
	fontData, err := truetype.Parse(gomono.TTF)
	if err != nil {
		panic(err)
	}

	titleFile, err := fs.Open("assets/pexels-ann-h-15009816-2.jpg")
	if err != nil {
		log.Fatalf("failed to find title background: %v", err)
	}

	defer titleFile.Close()

	titleBackground, err := io.ReadAll(titleFile)
	if err != nil {
		log.Fatalf("failed to find title background: %v", err)
	}

	titleScene := scene.NewTitle(titleBackground)

	mapTiles, err := maps.LoadTileImages(fs)
	if err != nil {
		log.Fatalf("failed to load map tiles: %v", err)
	}

	g := &Game{
		font:       truetype.NewFace(fontData, &truetype.Options{Size: 10}),
		titleScene: titleScene,
		fs:         fs,
		ecs:        ecs.NewECS(donburi.NewWorld()),
	}

	level := maps.NewLevel(g.Dimensions(), mapTiles)

	levelEntity := g.ecs.World.Entry(
		g.ecs.Create(
			layers.Floor,
			components.Level,
			components.Renderable,
		),
	)

	donburi.SetValue(
		levelEntity,
		components.Level,
		components.LevelComponent{
			Level: level,
		},
	)

	gameScene := &scene.GameScene{
		ECS: g.ecs,
	}

	g.gameScene = gameScene

	titleScene.Next(g.gameScreen)

	g.activeScene = g.titleScene

	g.ecs.
		AddSystem(system.Movement.Update).
		//AddRenderer(layers.Wall, system.Wall.Draw)
		AddRenderer(layers.Floor, system.LevelRender.Draw).
		AddRenderer(layers.RealWorld, system.Render.Draw)

	// g.ecs.
	// 	// AddSystem(system.NewSpawn().Update).
	// 	// AddSystem(metrics.Update).
	// 	// AddSystem(system.NewBounce(&g.bounds).Update).
	// 	// AddSystem(system.Velocity.Update).
	// 	// AddSystem(system.Gravity.Update).
	// 	AddRenderer(layers.LayerWall, system.DrawWall).
	// 	AddRenderer(layers.LayerFloor, system.DrawFloor)
	// 	//AddRenderer(layers.LayerMetrics, metrics.Draw).
	// 	//AddRenderer(layers.LayerBunnies, system.Render.Draw)

	return g
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		os.Exit(0)
	}

	// TODO this can generate more maps than we want
	// can we limit it?
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		// Load up a new map
		mapTiles, err := maps.LoadTileImages(g.fs)
		if err != nil {
			log.Fatalf("failed to load map tiles: %v", err)
		}

		level := maps.NewLevel(g.Dimensions(), mapTiles)

		// Update entity with new map data
		levelQuery := ecs.NewQuery(
			layers.Floor,
			filter.Contains(
				components.Level,
			),
		)

		levelEntity, ok := levelQuery.First(g.ecs.World)
		if !ok {
			return nil
		}

		donburi.SetValue(
			levelEntity,
			components.Level,
			components.LevelComponent{
				Level: level,
			},
		)

		// Update player location to spawn in a new room
		x, y := level.Rooms[0].Center()

		playerQuery := ecs.NewQuery(
			layers.RealWorld,
			filter.Contains(
				components.Player,
			),
		)

		playerEntity, ok := playerQuery.First(g.ecs.World)
		if !ok {
			return nil
		}

		donburi.SetValue(
			playerEntity,
			components.Position,
			components.PositionComponent{
				X: x,
				Y: y,
			},
		)
	}

	g.activeScene.Update()

	g.ecs.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.activeScene.Draw(screen)
}

func (g *Game) Layout(_, _ int) (int, int) {
	return 1080, 768
}

func (g *Game) titleScreen() {
	// Initialize Scene
	g.activeScene = g.titleScene
}

func (g *Game) gameScreen() {
	g.EnterWorld()
	g.activeScene = g.gameScene
}

func (g *Game) Dimensions() maps.Dimensions {
	w, h := g.Layout(0, 0)

	return maps.Dimensions{
		Width:      w / 16,
		Height:     h / 16,
		TileWidth:  16,
		TileHeight: 16,
	}
}

func (g *Game) loadAssetFromFS(path string) *ebiten.Image {
	f, err := g.fs.Open(path)
	if err != nil {
		log.Fatalf("failed to find asset: %v", err)
	}

	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatalf("failed to load asset: %v", err)
	}

	return ebiten.NewImageFromImage(img)
}
