package player

import "github.com/hajimehoshi/ebiten/v2"

func Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(Player.Obj.X, Player.Obj.Y)

	// changes sprite shown depending on what direction the player is going
	if Player.IsLeft {
		screen.DrawImage(Player.WalkingLeft[Player.WalkingStage], op)
	} else {
		screen.DrawImage(Player.WalkingRight[Player.WalkingStage], op)
	}
}
