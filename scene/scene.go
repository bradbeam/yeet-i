package scene

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Update()
	Draw(screen *ebiten.Image)
}
