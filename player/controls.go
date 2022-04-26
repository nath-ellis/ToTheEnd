package player

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func Controls() {
	Player.YSpeed += Player.Gravity // applies gravity

	// X Collision
	xSpeed := Player.XSpeed

	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if c := Player.Obj.Check(-xSpeed, 0, "tile"); c != nil { // left controls
			xSpeed = c.ContactWithCell(c.Cells[0]).X() // collides the player with the block
		}
		Player.IsWalking = true
		Player.Obj.X -= xSpeed // moves player
	} else if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) { // right controls
		if c := Player.Obj.Check(xSpeed, 0, "tile"); c != nil {
			xSpeed = c.ContactWithCell(c.Cells[0]).X() // collides the player with the block
		}
		Player.IsWalking = true
		Player.Obj.X += xSpeed // moves player
	} else {
		Player.IsWalking = false
	}

	// Jumping
	if !Player.Falling {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			Player.YSpeed = -Player.JumpSpeed
			Player.Falling = true
			Player.Jumping = true
		} else {
			Player.Jumping = false
		}
	}

	// Y Collision
	ySpeed := Player.YSpeed

	Player.Falling = true

	ySpeed = math.Max(math.Min(ySpeed, 32), -32)

	cd := ySpeed
	if ySpeed >= 0 {
		cd++
	}

	if c := Player.Obj.Check(0, cd, "tile"); c != nil {
		slide := c.SlideAgainstCell(c.Cells[0], "tile")

		if ySpeed < 0 && c.Cells[0].ContainsTags("tile") && slide != nil && math.Abs(slide.X()) <= 8 {
			Player.Obj.X += slide.X()
		} else {
			if tiles := c.ObjectsByTags("tile"); len(tiles) > 0 && (Player.Falling || Player.Obj.Y >= tiles[0].Y) {
				ySpeed = c.ContactWithObject(tiles[0]).Y()
				Player.YSpeed = 0

				// if you are on top of a tile
				if tiles[0].Y > Player.Obj.Y {
					Player.Falling = false
				}
			}
		}
	}

	Player.Obj.Y += ySpeed

	Player.Obj.Update()
}
