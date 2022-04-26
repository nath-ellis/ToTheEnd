package tiles

import (
	"github.com/nath-ellis/ToTheEnd/data"
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
		{resolv.NewObject(60, 536, 64, 64, "tile"), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64, "tile"), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64, "tile"), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64, "tile"), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64, "tile"), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64, "tile"), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64, "tile"), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64, "tile"), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64, "tile"), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64, "tile"), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64, "tile"), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64, "tile"), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64, "tile"), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64, "tile"), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64, "tile"), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64, "tile"), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64, "tile"), "TileGrass"},
		{resolv.NewObject(60, 536, 64, 64, "tile"), "Tile"},
		{resolv.NewObject(60, 536, 64, 64, "tile"), "Tile"},
		{resolv.NewObject(60, 536, 64, 64, "tile"), "Tile"},
		{resolv.NewObject(60, 536, 64, 64, "tile"), "Tile"},
		{resolv.NewObject(60, 536, 64, 64, "tile"), "Tile"},
		{resolv.NewObject(60, 536, 64, 64, "tile"), "Tile"},
		{resolv.NewObject(60, 536, 64, 64, "tile"), "Tile"},
		{resolv.NewObject(60, 536, 64, 64, "tile"), "Tile"},
		{resolv.NewObject(60, 536, 64, 64, "tile"), "Tile"},
	}

	// Adds them to the space
	for _, t := range Tiles {
		data.Space.Add(t.Obj)
	}

	// Randomizes the position of the tiles
	Randomize()
}
