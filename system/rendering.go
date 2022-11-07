package system

import (
	"math"

	"github.com/grzesl/ship/assets"
	"github.com/grzesl/ship/component"
	"github.com/grzesl/ship/helper/enum"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"
)

type Rendering struct {
	offscreen *ebiten.Image
}

func NewRendering(w, h int) *Rendering {
	return &Rendering{
		offscreen: ebiten.NewImage(w*assets.Tilesize, h*assets.Tilesize),
	}
}

func (t *Rendering) Draw(w engine.World, screen *ebiten.Image) {
	screen.Fill(assets.Background)

	// Get auxiliary objects
	cameraEntity, found := w.View(component.Pos{}, component.Zoom{}).Get()
	if !found {
		return
	}
	var camera *component.Pos
	var zoom *component.Zoom
	cameraEntity.Get(&camera, &zoom)

	if math.Abs(zoom.Value-zoom.DesiredZoomValue) < 0.001 {
		zoom.Value = zoom.DesiredZoomValue
	} else if zoom.Value > zoom.DesiredZoomValue {
		zoom.Value -= 0.001
	} else if zoom.Value < zoom.DesiredZoomValue {
		zoom.Value += 0.001
	}

	playerEntity, found := w.View(component.Pos{}, component.Rot{}, component.Control{}).Get()
	if !found {
		return
	}
	var pos *component.Pos
	var rot *component.Rot
	var control *component.Control
	playerEntity.Get(&pos, &rot, &control)

	// Draw tiles to the offscreen
	view := w.View(component.Solid{}, component.Pos{}, component.Sprite{})
	view.Each(func(e engine.Entity) {
		var solid *component.Solid
		var pos *component.Pos
		var sprite *component.Sprite
		var siz *component.Size
		e.Get(&solid, &pos, &sprite, &siz)

		//if solid.Empty() {
		//	return
		//}

		if e == playerEntity {
			zoom.DesiredZoomValue = 1 + (control.VolumeSpeed * 20)
		}

		if solid.Group == enum.CollisionGroupPlayer {
			return
		}

		//log.Println("Entityid:" + strconv.Itoa(e.ID()))

		op := &ebiten.DrawImageOptions{}

		if solid.Empty() {
			op.GeoM.Translate(pos.X*float64(assets.Tilesize),
				pos.Y*float64(assets.Tilesize))
		} else {
			op.GeoM.Translate(pos.X*float64(assets.Tilesize)-(float64(sprite.Frameset.Image().Bounds().Size().X)),
				pos.Y*float64(assets.Tilesize)-(float64(sprite.Frameset.Image().Bounds().Size().Y)))
		}

		t.offscreen.DrawImage(sprite.Frameset.Image(), op)
		//log.Println("sprite name:" + sprite.Name)
	})

	// Draw the offscreen
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(zoom.Value, zoom.Value)
	op.GeoM.Translate(-camera.X, -camera.Y)
	screen.DrawImage(t.offscreen, op)
	t.offscreen.Clear()
}
