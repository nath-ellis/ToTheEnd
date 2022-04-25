package tiles

import (
	"github.com/nath-ellis/ToTheEnd/space"
	"golang.org/x/exp/rand"
)

func Randomize() {
	for _, t := range Tiles {
		space.Space.Remove(t.Obj) // removes it

		t.Obj.X = float64(rand.Intn(960)) // randomizes x pos
		t.Obj.Y = float64(rand.Intn(536)) // randomizes y pos
		// Numbers are 64 less than screen width and screen height to make sure no tiles are off the screen

		space.Space.Add(t.Obj) // adds it back in the new position
	}
}
