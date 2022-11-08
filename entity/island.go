package entity

import "github.com/grzesl/ship/component"

// The object that takes up space for the tile.

type Island struct {
	component.Pos    // Index
	component.Vel    // Vel
	component.Size   // Current size in tiles
	component.Solid  // Solid or empty
	component.Sprite // Current image.
	component.Goods  // Island goods
}
