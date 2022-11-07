package system

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/grzesl/ship/component"
	"github.com/sedyh/mizu/pkg/engine"
)

type Render struct{}

func rotate_point(x float64, y float64, angle float64) (float64, float64) {
	s := math.Sin(angle)
	c := math.Cos(angle)

	xnew := x*c - y*s
	ynew := x*s + y*c

	return xnew, ynew
}

func (r *Render) Draw(w engine.World, screen *ebiten.Image) {
	//screen.Fill(assets.Background)

	// Get auxiliary objects
	cameraEntity, found := w.View(component.Pos{}, component.Zoom{}).Get()
	if !found {
		return
	}
	var camera *component.Pos
	var zoom *component.Zoom
	cameraEntity.Get(&camera, &zoom)

	playerEntity, found := w.View(component.Pos{}, component.Rot{}, component.Control{}).Get()
	if !found {
		return
	}
	var pos *component.Pos
	var rot *component.Rot
	var control *component.Control

	playerEntity.Get(&pos, &rot, &control)

	view := w.View(component.Root{}, component.Pivot{}, component.Pos{})
	// Get all particles
	view.Each(func(e engine.Entity) {
		var root *component.Root
		var pivot *component.Pivot
		var pos *component.Pos
		var angle *component.Angle
		var scale *component.Scale
		var life *component.Life
		var gradient *component.Gradient
		var sprite *component.Sprite
		e.Get(
			&root, &pivot, &pos, &angle,
			&scale, &life, &gradient, &sprite,
		)

		if root == nil { //not a particle?
			return
		}

		if root.Enabled {
			return
		}

		// Calculate parameters for image transformation.
		iw, ih := float64(sprite.Image.Bounds().Dx())*scale.W, float64(sprite.Image.Bounds().Dy())*scale.H
		sw, sh := float64(screen.Bounds().Dx()), float64(screen.Bounds().Dy())

		screenX, screenY := (pivot.X+pos.X)*sw, (pivot.Y+pos.Y)*sh

		theta := float64(angle.Deg) * math.Pi / 180

		age := float64(life.Current) / float64(life.Total)
		color := gradient.Colors[int(age*float64(len(gradient.Colors)-1))]
		cr, cg, cb, ca := color.RGBA()
		red, green, blue, alpha := float64(cr)/0xFFFF, float64(cg)/0xFFFF, float64(cb)/0xFFFF, float64(ca)/0xFFFF

		// Smooth fade in and out.
		startFadeIn := 0.1
		endFadeIn := 0.3
		startFadeOut := 0.5
		endFadeOut := 1.0
		if age <= endFadeIn {
			alpha = 0
			if age >= startFadeIn {
				alpha = age / (startFadeIn + endFadeIn)
			}
		} else if age >= startFadeOut {
			alpha = 1 - (age)/endFadeOut
		}

		// Apply camera zoom.
		//scaledTilesize := float64(assets.Tilesize) * zoom.Value

		// Draw the particle.
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(scale.W*zoom.Value*-control.VolumeSpeed*50,
			scale.H*zoom.Value*-control.VolumeSpeed*50)
		op.GeoM.Translate(-iw/2, -ih/2)
		op.GeoM.Rotate(theta)

		screenB := w.Bounds()

		op.ColorM.Scale(red, green, blue, alpha)

		screenX, screenY = rotate_point(screenX, screenY, rot.Radians)
		op.GeoM.Translate((float64(screenB.Dx())+screenX)/2-10, (float64(screenB.Dy())+screenY)/2-10)
		screen.DrawImage(sprite.Image, op)
	})
}
