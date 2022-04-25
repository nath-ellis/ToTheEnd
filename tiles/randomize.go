package tiles

import "golang.org/x/exp/rand"

func Randomize() {
	for _, t := range Tiles {
		t.Obj.X = float64(rand.Intn(960))
		t.Obj.Y = float64(rand.Intn(536))
	}
}
