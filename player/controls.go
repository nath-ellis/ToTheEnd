package player

import "github.com/hajimehoshi/ebiten/v2"

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
}
