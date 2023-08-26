package main

import (
	"bytes"
	"image"
	_ "image/jpeg"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Title struct {
	background *ebiten.Image
	next       func()
}

func NewTitle(background []byte) *Title {
	img, _, err := image.Decode(bytes.NewReader(background))
	if err != nil {
		log.Fatalf("failed to load background image: %v", err)
	}

	return &Title{
		background: ebiten.NewImageFromImage(img),
	}
}

func (t *Title) Next(next func()) {
	t.next = next
}

func (t *Title) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) ||
		ebiten.IsKeyPressed(ebiten.KeySpace) ||
		ebiten.IsKeyPressed(ebiten.KeyEscape) ||
		ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) ||
		ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		t.next()
	}

}

func (t *Title) Draw(screen *ebiten.Image) {
	screen.DrawImage(t.background, &ebiten.DrawImageOptions{})
}
