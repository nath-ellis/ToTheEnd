package space

import "github.com/solarlune/resolv"

// The space is in it's own file so it can be accessed anywhere easily
var Space *resolv.Space = resolv.NewSpace(1024, 600, 1, 1)
