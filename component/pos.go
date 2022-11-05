package component

import "github.com/grzesl/ship/helper/random"

type Pos struct {
	X, Y float64
	Init func() Pos
}

func NewPosI(x, y int) Pos {
	return Pos{float64(x), float64(y), nil}
}

func NewPos(xMin, xMax, yMin, yMax float64) Pos {
	init := func() Pos {
		return Pos{
			X: random.RangeFloat(xMin, xMax),
			Y: random.RangeFloat(yMin, yMax),
		}
	}
	res := init()
	res.Init = init
	return res
}
