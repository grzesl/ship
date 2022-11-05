package component

import (
	"github.com/grzesl/ship/helper/random"
)

type Life struct {
	Current, Total int
	Init           func() Life
}

func NewLife(currentMin, currentMax, totalMin, totalMax int) Life {
	init := func() Life {
		return Life{
			Current: random.RangeInt(currentMin, currentMax),
			Total:   random.RangeInt(totalMin, totalMax),
		}
	}
	res := init()
	res.Init = init
	return res
}
