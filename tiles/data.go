package tiles

import (
	"github.com/nath-ellis/ToTheEnd/space"
	"github.com/solarlune/resolv"
)

type Tile struct {
	Obj  *resolv.Object
	Type string
}

var (
	Tiles []Tile
)

func Init() {
	Tiles = []Tile{
		{resolv.NewObject(60, 536, 64, 64), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64), "Tile"},
		{resolv.NewObject(60, 536, 64, 64), "Tile"},
		{resolv.NewObject(60, 536, 64, 64), "Tile"},
		{resolv.NewObject(60, 536, 64, 64), "Tile"},
		{resolv.NewObject(60, 536, 64, 64), "Tile"},
		{resolv.NewObject(60, 536, 64, 64), "Tile"},
		{resolv.NewObject(60, 536, 64, 64), "Tile"},
		{resolv.NewObject(60, 536, 64, 64), "Tile"},
		{resolv.NewObject(60, 536, 64, 64), "Tile"},
	}

	for _, t := range Tiles {
		space.Space.Add(t.Obj)
	}

	Randomize()
}
