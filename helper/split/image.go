package split

import (
	"time"

	"github.com/grzesl/ship/helper/geo"
	"github.com/grzesl/ship/helper/graphics"
	"github.com/hajimehoshi/ebiten/v2"
)

// Single cuts one image from the spritesheet and mirrors it if necessary.
func Single(source *ebiten.Image, x, y, w, h int, flipped bool) *graphics.Frameset {
	result := source.SubImage(geo.Rect(x, y, w, h)).(*ebiten.Image)
	if flipped {
		result = flip(result)
	}
	return graphics.NewFramesetSingle(result)
}

// Multi cuts several images from the spritesheet along a line and mirrors it if necessary.
func Multi(source *ebiten.Image, x, y, w, h, count int, flipped bool, speed time.Duration) *graphics.Frameset {
	var result []*ebiten.Image
	for i := 0; i < count; i++ {
		img := source.SubImage(geo.Rect(x+w*i, y, w, h)).(*ebiten.Image)
		if flipped {
			img = flip(img)
		}
		result = append(result, img)
	}
	return graphics.NewFramesetMulti(speed, result...)
}

// Multi cuts several images from the spritesheet along a line and mirrors it if necessary.
func MultiRows(source *ebiten.Image, x, y, w, h, cols, count int, flipped bool, speed time.Duration) *graphics.Frameset {
	var result []*ebiten.Image
	for i := 0; i < count; i++ {

		row := i / cols
		col := i % cols

		img := source.SubImage(geo.Rect(w*col, h*row, w, h)).(*ebiten.Image)
		if flipped {
			img = flip(img)
		}
		result = append(result, img)
	}
	return graphics.NewFramesetMulti(speed, result...)
}

func flip(source *ebiten.Image) *ebiten.Image {
	img := ebiten.NewImage(source.Size())
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(-1, 1)
	op.GeoM.Translate(float64(img.Bounds().Dx()), 0)
	img.DrawImage(source, op)
	return img
}
