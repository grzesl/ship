package component

import (
	"github.com/grzesl/ship/helper/random"
)

type Scale struct {
	W, H float64
	Init func() Scale
}

func NewScale(min, max float64) Scale {
	init := func() Scale {
		wh := random.RangeFloat(min, max)

		return Scale{
			W: wh,
			H: wh,
		}
	}
	res := init()
	res.Init = init
	return res
}
