package main

import (
	"embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/bradbeam/yeet-i/game"
)

//go:embed assets
var assetsFS embed.FS

func main() {
	ebiten.SetWindowSize(1080, 768)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(game.NewGame(assetsFS)); err != nil {
		log.Fatal(err)
	}
}
