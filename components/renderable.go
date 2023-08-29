package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type RenderableComponent struct {
	Image *ebiten.Image
}

var Renderable = donburi.NewComponentType[RenderableComponent]()
