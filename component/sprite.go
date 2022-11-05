package component

import (
	"github.com/grzesl/ship/helper/graphics"
	"github.com/grzesl/ship/helper/random"
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Name     string
	Frameset *graphics.Frameset
	Image    *ebiten.Image
	Init     func() Sprite
}

func NewSpriteFs(name string, frameset *graphics.Frameset) Sprite {
	return Sprite{Name: name, Frameset: frameset}
}

func NewSprite(images ...*ebiten.Image) Sprite {
	init := func() Sprite {
		index := random.RangeInt(0, len(images))
		return Sprite{Image: images[index]}
	}
	res := init()
	res.Init = init
	return res
}
