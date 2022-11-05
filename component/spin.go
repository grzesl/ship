package component

import (
	"github.com/grzesl/ship/helper/random"
)

type Spin struct {
	Speed int
	Init  func() Spin
}

func NewSpin(min, max int) Spin {
	init := func() Spin {
		return Spin{Speed: random.RangeInt(min, max)}
	}
	res := init()
	res.Init = init
	return res
}
