package system

import (
	"github.com/grzesl/ship/component"
	"github.com/sedyh/mizu/pkg/engine"
)

type Rotation struct {
	*component.Rot
}

func NewRotation() *Rotation {
	return &Rotation{}
}

func (v *Rotation) Update(_ engine.World) {
	// Increase position.
	//v.Pos.X += v.Vel.L
	//v.Pos.Y += v.Vel.M
}
