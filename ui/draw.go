package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	randomize, _, _ = ebitenutil.NewImageFromFile("assets/icons/randomize.png")
)

func Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(940, 268)

	screen.DrawImage(randomize, op)
}
