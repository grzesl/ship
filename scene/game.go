package scene

import (
	"math/rand"
	"time"

	"github.com/grzesl/ship/assets"
	"github.com/grzesl/ship/component"
	"github.com/grzesl/ship/entity"
	"github.com/grzesl/ship/helper/enum"
	"github.com/grzesl/ship/helper/load"
	"github.com/grzesl/ship/helper/tilemap"
	"github.com/grzesl/ship/system"
	"github.com/sedyh/mizu/pkg/engine"
)

type Game struct {
}

func (g *Game) Setup(w engine.World) {
	rand.Seed(time.Now().UTC().UnixNano())
	assets.Init()

	w.AddComponents(
		component.Root{}, component.Birthrate{}, component.Pivot{},
		component.Pos{}, component.Vel{}, component.Size{}, component.Rot{}, component.Zoom{},
		component.Solid{}, component.Construct{}, component.Sprite{}, component.Gravity{}, component.Control{},
		component.Accel{}, component.Spin{}, component.Scale{}, component.Angle{}, component.Growth{},
		component.Life{}, component.Gradient{},
	)

	// Fills the entire space with tiles, some of them are marked as empty.
	construct := component.NewConstruct(load.Level(assets.Level))
	w.AddSystems(
		system.NewRendering(construct.Width, construct.Height),
		&system.Emit{}, &system.Acceleration{}, &system.Spin{}, &system.Growth{}, &system.Age{},
		&system.Death{}, &system.Render{},
		system.NewShip(), system.NewGravity(), system.NewCollision(), system.NewVelocity(),
		system.NewFocus(), system.NewConnection(), system.NewDebug(), system.NewAnimation(),
	)

	w.AddEntities(&entity.Ocean{
		Pos: component.NewPosI(0, 0),
		Size: component.NewSizeF(
			float64(assets.OceanH)/float64(assets.Tilesize),
			float64(assets.OceanH)/float64(assets.Tilesize),
		),
		Solid:  component.NewSolid(enum.CollisionGroupNone),
		Sprite: component.NewSpriteFs("ocean", assets.Images["ocean_idle"]),
	})

	for y := 0; y < construct.Height; y++ {
		for x := 0; x < construct.Width; x++ {
			switch string(construct.Level[tilemap.Index(x, y, construct.Width)]) {
			case "1":
				w.AddEntities(&entity.Tile{
					Pos:    component.NewPosI(x, y),
					Vel:    component.NewVelF(0, 0),
					Size:   component.NewSizeI(1, 1),
					Solid:  component.NewSolid(enum.CollisionGroupTile),
					Sprite: component.NewSpriteFs("rocks", assets.Images["rocks_idle"]),
				})
				/*default:
				w.AddEntities(&entity.Tile{
					Pos:    component.NewPosI(x, y),
					Vel:    component.NewVel(0, 0),
					Size:   component.NewSizeI(1, 1),
					Solid:  component.NewSolid(enum.CollisionGroupNone),
					Sprite: component.NewSprite("wall", assets.Images["wall_0"]),
				})*/
			}

		}
	}

	/*
		for y := 0; y < construct.Height; y++ {
			for x := 0; x < construct.Width; x++ {
				switch string(construct.Level[tilemap.Index(x, y, construct.Width)]) {
				case "o":
					w.AddEntities(&entity.Tile{
						Pos:    component.NewPosI(x, y),
						Vel:    component.NewVel(0, 0),
						Size:   component.NewSizeI(1, 1),
						Solid:  component.NewSolid(enum.CollisionGroupTile),
						Sprite: component.NewSprite("wall", assets.Images["wall_0"]),
					})
				default:
					w.AddEntities(&entity.Tile{
						Pos:    component.NewPosI(x, y),
						Vel:    component.NewVel(0, 0),
						Size:   component.NewSizeI(1, 1),
						Solid:  component.NewSolid(enum.CollisionGroupNone),
						Sprite: component.NewSprite("wall", assets.Images["wall_0"]),
					})
				}

			}
		}
	*/
	// Adding crates.
	/*
		for y := 0; y < construct.Height; y++ {
			for x := 0; x < construct.Width; x++ {
				switch string(construct.Level[tilemap.Index(x, y, construct.Width)]) {
				case "x":
					w.AddEntities(&entity.Crate{
						Pos: component.NewPosI(x, y),
						Size: component.NewSizeF(
							float64(assets.Cratesize)/float64(assets.Tilesize),
							float64(assets.Cratesize)/float64(assets.Tilesize),
						),
						Vel:     component.NewVel(0, 0),
						Solid:   component.NewSolid(enum.CollisionGroupCrate),
						Sprite:  component.NewSprite("crate", assets.Images["crate"]),
						Gravity: component.NewGravity(0.04),
					})
				}
			}
		}*/

	// Adding a player.
	/*w.AddEntities(
		&entity.Player{
			Pos: component.NewPosI(construct.Width/2, construct.Height/2),
			Size: component.NewSizeF(
				float64(assets.PlayerW)/float64(assets.Tilesize),
				float64(assets.PlayerH)/float64(assets.Tilesize),
			),
			Vel:     component.NewVel(0, 0),
			Solid:   component.NewSolid(enum.CollisionGroupPlayer),
			Sprite:  component.NewSprite(assets.Images["player_run_right"]),
			Gravity: component.NewGravity(0.04),
			Control: component.NewControl(0.7, 2.5, 2.0, 0.2),
		},
		&entity.Construct{Construct: construct}, &entity.Camera{Zoom: component.NewZoom(1.0)},
	)*/

	w.AddEntities(
		&entity.Player{
			Pos: component.NewPosI(construct.Width/2, construct.Height/2),
			Size: component.NewSizeF(
				float64(assets.ShipW)/float64(assets.Tilesize),
				float64(assets.ShipH)/float64(assets.Tilesize),
			),
			Vel:     component.NewVelF(0, 0),
			Solid:   component.NewSolid(enum.CollisionGroupPlayer),
			Sprite:  component.NewSpriteFs("ship", assets.Images["ship_idle"]),
			Gravity: component.NewGravity(0.00),
			Control: component.NewControl(0.7, 2.5, 2.0, 0.001),
			Rot:     component.NewRotF(0.0),
		},
		&entity.Construct{Construct: construct}, &entity.Camera{Zoom: component.NewZoom(1.0)},
	)

	w.AddEntities(
		&entity.Emitter{
			Root:      component.NewRoot(),
			Birthrate: component.NewBirthrate(2, 20*time.Millisecond),
			Pivot:     component.NewPivot(0.0, 0.0),
			Pos:       component.NewPos(0.00, 0.00, 0.0, 0.0),
			Vel:       component.NewVel(0.01, 0.01, 80, 100),
			Accel:     component.NewAccel(-0.0001, -0.0001, 40, 90),
			Angle:     component.NewAngle(-90, 90),
			Spin:      component.NewSpin(1, 2),
			Scale:     component.NewScale(0.25, 0.25),
			Growth:    component.NewGrowth(-0.007, -0.003),
			Life:      component.NewLife(0, 10, 30, 40),
			Gradient:  component.NewGradient(assets.WaterGradient...),
			Sprite:    component.NewSprite(assets.Water),
		})
}
