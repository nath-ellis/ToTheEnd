package tiles

import (
	"github.com/nath-ellis/ToTheEnd/data"
	"golang.org/x/exp/rand"
)

func Randomize() {
	for _, t := range Tiles {
		data.Space.Remove(t.Obj) // removes it

		// Randomizes the position of the tiles
		t.Obj.X = float64(rand.Intn(960))
		t.Obj.Y = float64(rand.Intn(536))

		data.Space.Add(t.Obj) // adds it back in the new position
	}
}
