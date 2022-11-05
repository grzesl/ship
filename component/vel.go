package component

import (
	"math"

	"github.com/grzesl/ship/helper/random"
)

type Vel struct {
	L, M float64
	Init func() Vel
}

func NewVelF(l, m float64) Vel {
	return Vel{l, m, nil}
}

func NewVel(speedMin, speedMax, directionMin, directionMax float64) Vel {
	init := func() Vel {
		speed := random.RangeFloat(speedMin, speedMax)
		direction := random.RangeFloat(directionMin, directionMax)

		return Vel{
			L: math.Cos(-direction*math.Pi/180) * speed,
			M: math.Sin(-direction*math.Pi/180) * speed,
		}
	}
	res := init()
	res.Init = init
	return res
}
