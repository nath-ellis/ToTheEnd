package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nath-ellis/ToTheEnd/tiles"
)

var btnPressCooldown int = 0 // cooldown so you cannot hold down the button

func Update() {
	mouseX, mouseY := ebiten.CursorPosition()

	mouseObj.X = float64(mouseX)
	mouseObj.Y = float64(mouseY)

	mouseObj.Update()

	if btnPressCooldown <= 0 {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			// if the cursor is on the button
			if c := mouseObj.Check(0, 0, "randomize"); c != nil {
				tiles.Randomize()
				btnPressCooldown += 50
			}
		}
	} else {
		btnPressCooldown -= 1
	}
}
