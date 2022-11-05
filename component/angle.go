package component

import (
	"github.com/grzesl/ship/helper/random"
)

type Angle struct {
	Deg  int
	Init func() Angle
}

func NewAngle(min, max int) Angle {
	init := func() Angle {
		return Angle{Deg: random.RangeInt(min, max)}
	}
	res := init()
	res.Init = init
	return res
}
