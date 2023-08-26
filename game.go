package main

import (
	"embed"
	"log"
	"os"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomono"
)

type Scene interface {
	Update()
	Draw(screen *ebiten.Image)
}

type Game struct {
	activeScene Scene
	font        font.Face

	titleScene Scene
	gameScene  Scene
}

func NewGame(fs embed.FS) *Game {
	fontData, err := truetype.Parse(gomono.TTF)
	if err != nil {
		panic(err)
	}

	titleBackground, err := fs.ReadFile("assets/pexels-ann-h-15009816-2.jpg")
	if err != nil {
		log.Fatalf("failed to find title background: %v", err)
	}
	titleScene := NewTitle(titleBackground)

	g := &Game{
		font:       truetype.NewFace(fontData, &truetype.Options{Size: 10}),
		titleScene: titleScene,
		gameScene:  &GameScene{},
	}

	g.activeScene = g.titleScene

	titleScene.Next(g.gameScreen)

	return g
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		os.Exit(0)
	}

	g.activeScene.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.activeScene.Draw(screen)
}

func (g *Game) Layout(width, height int) (int, int) {
	return width, height
}

func (g *Game) titleScreen() {
	// Initialize Scene
	g.activeScene = g.titleScene
}

func (g *Game) gameScreen() {
	g.activeScene = g.gameScene
}
