package system

import (
	"math"

	"github.com/grzesl/ship/helper/num"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/grzesl/ship/assets"
	"github.com/grzesl/ship/component"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"
)

type Ship struct{}

func NewShip() *Ship {
	return &Ship{}
}

func (p *Ship) Update(w engine.World) {
	// Get auxiliary objects
	player, ok := w.View(component.Pos{}, component.Vel{}, component.Sprite{}, component.Control{}, component.Rot{}).Get()
	if !ok {
		return
	}
	var pos *component.Pos
	var vel *component.Vel
	var rot *component.Rot
	var gravity *component.Gravity
	var sprite *component.Sprite
	var control *component.Control
	player.Get(&pos, &vel, &gravity, &sprite, &control, &rot)

	// Almost a jump buffer.
	grounded := num.Equal(vel.M, 0, 0.002)

	// Jump with height control.
	if grounded && inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		vel.M -= control.JumpSpeed

	}
	if vel.M > 0 {
		vel.M += gravity.Value * (control.FallSpeed - 1)

	} else if vel.M < 0 && !ebiten.IsKeyPressed(ebiten.KeySpace) {
		vel.M += gravity.Value * (control.LowSpeed - 1)

	}

	// Horizontal movement.
	//moveDirection := 0.0
	//moveForward := 0.0

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		//moveDirection = -1.0
		rot.Radians -= 0.01
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		//moveDirection = 1.0
		rot.Radians += 0.01
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		//moveForward = -1.0
		//rot.Radius -= 0.01
		control.VolumeSpeed -= control.MoveSpeed

		if control.VolumeSpeed < -0.02 {
			control.VolumeSpeed = -0.02
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		//moveForward = 1.0
		//rot.Radius += 0.01
		control.VolumeSpeed += 0.0001 //break
		if control.VolumeSpeed > 0 {
			control.VolumeSpeed = 0
		}
	} else {

		control.VolumeSpeed += 0.00001 //slow down
		if control.VolumeSpeed > 0 {
			control.VolumeSpeed = 0
		}
	}

	//log.Println("F: ", moveForward, " D: ", moveDirection)
	//log.Println("L: ", vel.L, " M: ", vel.M, " S: ", control.MoveSpeed)

	vel.M = -(math.Cos(rot.Radians) * control.VolumeSpeed) // X-component.
	vel.L = (math.Sin(rot.Radians) * control.VolumeSpeed)  // Y-component

	//vel.L = num.Lerp(vel.L, control.MoveSpeed*moveDirection, control.MoveSpeed)
	//vel.M = num.Lerp(vel.M, control.MoveSpeed*moveForward, control.MoveSpeed)
	// Player movement animation.
	// You can easily add sharp turns here if you make the horizontal movement smoother.
	//if moveDirection != 0 {
	//	sprite.Frameset = assets.Images["player_"+name.PlayerMovement(grounded)+"_"+name.PlayerDirection(moveDirection)]
	//} else {
	//		sprite.Frameset = assets.Images["player_idle"]
	//	}

	sprite.Frameset = assets.Images["ship_idle"]
}

func (p *Ship) Draw(w engine.World, screen *ebiten.Image) {
	// Get auxiliary objects
	player, ok := w.View(component.Pos{}, component.Vel{}, component.Sprite{}, component.Rot{}).Get()
	if !ok {
		return
	}
	cameraEntity, found := w.View(component.Pos{}, component.Zoom{}).Get()
	if !found {
		return
	}
	var pos *component.Pos
	var vel *component.Vel
	var rot *component.Rot
	var sprite *component.Sprite
	var camera *component.Pos
	var zoom *component.Zoom
	player.Get(&pos, &vel, &sprite, &rot)
	cameraEntity.Get(&camera, &zoom)

	// Apply camera zoom.
	scaledTilesize := float64(assets.Tilesize) * zoom.Value

	// Draw the player.
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(-float64(assets.ShipW)/2, -float64(assets.ShipH)/2)

	op.GeoM.Scale(zoom.Value, zoom.Value)
	op.GeoM.Rotate(rot.Radians)

	//op.GeoM.Translate(float64(assets.ShipW)/2, float64(assets.ShipH)/2)
	op.GeoM.Translate(-camera.X, -camera.Y)
	op.GeoM.Translate(pos.X*scaledTilesize, pos.Y*scaledTilesize)

	screen.DrawImage(sprite.Frameset.Image(), op)
}
