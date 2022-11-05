package system

import (
	"github.com/grzesl/ship/component"
	"github.com/sedyh/mizu/pkg/engine"
)

type Animation struct {
	*component.Sprite
}

func NewAnimation() *Animation {
	return &Animation{}
}

func (a *Animation) Update(_ engine.World) {

	if a.Frameset == nil {
		return //gl
	}
	// Sprites containing one frame show only it.
	if len(a.Frameset.Images) < 2 {
		return
	}

	// In sprites containing multiple frames, the current frame is constantly changing at the specified interval.
	select {
	case <-a.Frameset.Ticker.C:
		a.Frameset.Current = (a.Frameset.Current + 1) % len(a.Frameset.Images)
	default:
	}
}
