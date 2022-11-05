package component

import (
	"github.com/grzesl/ship/helper/random"
)

type Growth struct {
	Speed float64
	Init  func() Growth
}

func NewGrowth(min, max float64) Growth {
	init := func() Growth {
		return Growth{Speed: random.RangeFloat(min, max)}
	}
	res := init()
	res.Init = init
	return res
}
