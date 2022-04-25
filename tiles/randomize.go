package tiles

import (
	"math/rand"

	"github.com/nath-ellis/ToTheEnd/data"
)

func Randomize() {
	chance := rand.Intn(5)
	collectiveYPos := rand.Intn(data.SCREEN_HEIGHT - 64)

	for _, t := range Tiles {
		data.Space.Remove(t.Obj) // removes it

		// Randomizes the position of the tiles
		// Numbers are 64 less than screen width and screen height to make sure no tiles are off the screen
		if chance <= 1 {
			// Completely random
			t.Obj.X = float64(rand.Intn(data.SCREEN_WIDTH - 64))
			t.Obj.Y = float64(rand.Intn(data.SCREEN_HEIGHT - 64))
		} else if chance == 2 {
			// All in a line but some in random places
			t.Obj.X = float64(rand.Intn(data.SCREEN_WIDTH - 64))

			chance2 := rand.Intn(10)
			if (chance2 % 2) == 0 {
				t.Obj.Y = float64(collectiveYPos)
			} else {
				t.Obj.Y = float64(data.SCREEN_HEIGHT - 64)
			}
		} else if chance == 3 {
			// Spread apart from a line
			t.Obj.X = float64(rand.Intn(data.SCREEN_WIDTH - 64))

			amount := rand.Intn(256)
			chance2 := rand.Intn(2)

			if chance2 <= 1 {
				t.Obj.Y = float64(collectiveYPos - amount)
			} else {
				t.Obj.Y = float64(collectiveYPos + amount)
			}
		} else if chance == 4 {
			// All completely in a line
			t.Obj.X = float64(rand.Intn(data.SCREEN_WIDTH - 64))
			t.Obj.Y = float64(collectiveYPos)
		}

		data.Space.Add(t.Obj) // adds it back in the new position
	}
}
