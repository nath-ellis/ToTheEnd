package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nath-ellis/ToTheEnd/data"
	"github.com/solarlune/resolv"
)

var mouseObj *resolv.Object

func Init() {
	mouseX, mouseY := ebiten.CursorPosition()

	mouseObj = resolv.NewObject(float64(mouseX), float64(mouseY), 10, 10, "mouse")

	data.Space.Add(mouseObj)
}
