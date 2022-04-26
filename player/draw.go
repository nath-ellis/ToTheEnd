package player

import "github.com/hajimehoshi/ebiten/v2"

// temporary for testing the player controls
func Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(Player.Obj.X, Player.Obj.Y)

	screen.DrawImage(Player.WalkingLeft[0], op)
}
