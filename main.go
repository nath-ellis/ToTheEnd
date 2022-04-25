package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nath-ellis/ToTheEnd/data"
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
	return data.SCREEN_WIDTH, data.SCREEN_HEIGHT
}

func main() {
	ebiten.SetWindowSize(data.SCREEN_WIDTH, data.SCREEN_HEIGHT)
	ebiten.SetWindowTitle("To The End")

	// Runs the game
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal("Error Running the Game: ", err)
	}
}
