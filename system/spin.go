package system

import (
	"github.com/grzesl/ship/component"
	"github.com/sedyh/mizu/pkg/engine"
)

type Spin struct {
	*component.Root
	*component.Angle
	*component.Spin
}

func (s *Spin) Update(_ engine.World) {
	if s.Root.Enabled {
		return
	}

	// Increase rotation angle.
	s.Angle.Deg += s.Spin.Speed
}
