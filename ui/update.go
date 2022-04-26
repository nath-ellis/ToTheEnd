package ui

import "github.com/hajimehoshi/ebiten/v2"

func Update() {
	mouseX, mouseY := ebiten.CursorPosition()

	mouseObj.X = float64(mouseX)
	mouseObj.Y = float64(mouseY)

	mouseObj.Update()
}
