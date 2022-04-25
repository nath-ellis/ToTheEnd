package data

import "github.com/solarlune/resolv"

// Data package is used for data that needs to be accessed everywhere

var (
	Space *resolv.Space = resolv.NewSpace(1024, 600, 1, 1)
)

const (
	SCREEN_WIDTH  int = 1024
	SCREEN_HEIGHT int = 600
)
