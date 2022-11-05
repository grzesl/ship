package entity

import "github.com/grzesl/ship/component"

type Ocean struct {
	component.Pos    // Offset for rendering.
	component.Solid  // Collision group.
	component.Size   // Current size in tiles.
	component.Sprite // Current image.
}
