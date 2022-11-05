package component

import (
	"math"

	"github.com/grzesl/ship/helper/random"
)

type Accel struct {
	O, P float64
	Init func() Accel
}

func NewAccel(speedMin, speedMax, directionMin, directionMax float64) Accel {
	init := func() Accel {
		speed := random.RangeFloat(speedMin, speedMax)
		direction := random.RangeFloat(directionMin, directionMax)

		return Accel{
			O: math.Cos(-direction*math.Pi/180) * speed,
			P: math.Sin(-direction*math.Pi/180) * speed,
		}
	}
	res := init()
	res.Init = init
	return res
}
