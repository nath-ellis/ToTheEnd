package tiles

import "github.com/solarlune/resolv"

type Tile struct {
	Obj  *resolv.Object
	Type string
}

var (
	Tiles []Tile
)

func Init() {
	Tiles = []Tile{
		{resolv.NewObject(60, 80, 64, 64), "Tile"},
	}
}
