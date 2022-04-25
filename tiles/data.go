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
	// All of the tiles are in array to make it easier to interact with them
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

	// Adds them to the space
	for _, t := range Tiles {
		space.Space.Add(t.Obj)
	}

	// Randomizes the position of the tiles
	Randomize()
}
