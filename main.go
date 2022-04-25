package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nath-ellis/ToTheEnd/tiles"
)

func init() {
	tiles.Init()
}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	tiles.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1024, 600
}

func main() {
	ebiten.SetWindowSize(1024, 600)
	ebiten.SetWindowTitle("To The End")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
