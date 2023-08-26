package main

import (
	"embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets
var assetsFS embed.FS

func main() {
	ebiten.SetWindowSize(1080, 768)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(NewGame(assetsFS)); err != nil {
		log.Fatal(err)
	}
}
