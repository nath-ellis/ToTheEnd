package tiles

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	TileDefault, _, _ = ebitenutil.NewImageFromFile("assets/tiles/tile.png")
	TileGrass, _, _   = ebitenutil.NewImageFromFile("assets/tiles/tile_grass.png")
)

func Draw(screen *ebiten.Image) {
	for _, t := range Tiles {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(t.Obj.X, t.Obj.Y) // sets the position

		// Draws the sprite based on the type variable
		switch t.Type {
		case "Tile":
			screen.DrawImage(TileDefault, op)
		case "TileGrass":
			screen.DrawImage(TileGrass, op)
		}
	}
}
