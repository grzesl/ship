package component

import "time"

type Goods struct {
	IncType int
	OutType int
	Visited time.Time
}

func NewGoods() Goods {
	return Goods{0, 0, time.Now().Add(-time.Minute)}
}
