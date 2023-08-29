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
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomono"

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

	gameScene := &scene.GameScene{
		Level: maps.NewLevel(g.Dimensions(), mapTiles),
	}

	g.gameScene = gameScene

	titleScene.Next(g.gameScreen)

	g.activeScene = g.titleScene

	g.ecs.
		AddSystem(system.Movement.Update).
		//AddRenderer(layers.Wall, system.Wall.Draw)
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
		mapTiles, err := maps.LoadTileImages(g.fs)
		if err != nil {
			log.Fatalf("failed to load map tiles: %v", err)
		}

		gameScene := &scene.GameScene{
			Level: maps.NewLevel(g.Dimensions(), mapTiles),
		}

		g.activeScene = gameScene
	}

	g.activeScene.Update()

	g.ecs.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.activeScene.Draw(screen)

	for _, layer := range []ecs.LayerID{layers.Default, layers.Floor, layers.Wall, layers.RealWorld} {
		g.ecs.DrawLayer(layer, screen)
	}
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
