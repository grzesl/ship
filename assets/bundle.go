package assets

import (
	"embed"
	"image/color"
	"time"

	"github.com/grzesl/ship/helper"
	"github.com/grzesl/ship/helper/graphics"

	"github.com/grzesl/ship/helper/load"
	"github.com/grzesl/ship/helper/split"
)

// This is where the images, colors and levels are loaded.

const (
	Tilesize = 32
	PlayerW  = 25
	PlayerH  = 36

	ShipW = 80
	ShipH = 128

	OceanW = 1920
	OceanH = 1080

	RocksW = 104
	RocksH = 91

	Cratesize = 64

	Animation = 100 * time.Millisecond

	Level = `
		oooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo
		o                                                          o
		1                                                          o
		o        1                                       2         o
		1                                                          o
		o                                                          o
		1                                                          o
		o                                                          o
		1                                                          o
		o                                                          o
		1                                                          o
		o                                                          o
		1                                                          o
		o                                                          o
		1                                                          o
		o                                                          o
		1                                                          o
		o                                                          o
		1                                                          o
		o                                                          o
		1                                                          o
		o                                                          o
		1                                                          o
		o              2                               2           o
		1                                                          o
		o                                                          o
		1                                                          o
		o                                                          o
		1                                                          o
		o                                                          o
		1                                                          o
		o                                                          o
		oooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo
	`
)

var (
	Background = color.RGBA{R: 98, G: 164, B: 197, A: 255}
	Images     = make(map[string]*graphics.Frameset)
	//FireA        = helper.Image(fs, "data/image/fire-a.png")
	//FireB        = helper.Image(fs, "data/image/fire-b.png")
	Ebitengine = helper.Image(fs, "data/image/ebitengine.png")
	Water      = helper.Image(fs, "data/image/water.png")
	/*FireGradient = []color.Color{
		color.RGBA{R: 23, G: 13, B: 4, A: 255},
		color.RGBA{R: 36, G: 18, B: 6, A: 255},
		color.RGBA{R: 49, G: 24, B: 8, A: 255},
		color.RGBA{R: 63, G: 30, B: 10, A: 255},
		color.RGBA{R: 76, G: 35, B: 12, A: 255},
		color.RGBA{R: 90, G: 41, B: 14, A: 255},
		color.RGBA{R: 103, G: 47, B: 16, A: 255},
		color.RGBA{R: 117, G: 52, B: 18, A: 255},
		color.RGBA{R: 130, G: 58, B: 20, A: 255},
		color.RGBA{R: 143, G: 64, B: 22, A: 255},
		color.RGBA{R: 157, G: 69, B: 24, A: 255},
		color.RGBA{R: 170, G: 75, B: 26, A: 255},
		color.RGBA{R: 184, G: 81, B: 28, A: 255},
		color.RGBA{R: 193, G: 85, B: 31, A: 255},
		color.RGBA{R: 198, G: 87, B: 33, A: 255},
		color.RGBA{R: 203, G: 89, B: 35, A: 255},
		color.RGBA{R: 208, G: 92, B: 37, A: 255},
		color.RGBA{R: 213, G: 94, B: 39, A: 255},
		color.RGBA{R: 218, G: 97, B: 41, A: 255},
		color.RGBA{R: 223, G: 99, B: 44, A: 255},
		color.RGBA{R: 228, G: 102, B: 46, A: 255},
		color.RGBA{R: 233, G: 104, B: 48, A: 255},
		color.RGBA{R: 238, G: 106, B: 50, A: 255},
		color.RGBA{R: 243, G: 109, B: 52, A: 255},
		color.RGBA{R: 248, G: 111, B: 54, A: 255},
		color.RGBA{R: 253, G: 114, B: 57, A: 255},
		color.RGBA{R: 253, G: 114, B: 57, A: 255},
		color.RGBA{R: 253, G: 114, B: 57, A: 255},
		color.RGBA{R: 253, G: 114, B: 57, A: 255},
		color.RGBA{R: 253, G: 114, B: 57, A: 255},
		color.RGBA{R: 253, G: 114, B: 57, A: 255},
		color.RGBA{R: 253, G: 114, B: 57, A: 255},
		color.RGBA{R: 253, G: 114, B: 57, A: 255},
		color.RGBA{R: 253, G: 114, B: 57, A: 255},
		color.RGBA{R: 253, G: 114, B: 57, A: 255},
		color.RGBA{R: 253, G: 114, B: 57, A: 255},
		color.RGBA{R: 253, G: 114, B: 57, A: 255},
		color.RGBA{R: 253, G: 114, B: 57, A: 255},
		color.RGBA{R: 253, G: 114, B: 57, A: 255},
		color.RGBA{R: 253, G: 114, B: 57, A: 255},
		color.RGBA{R: 253, G: 114, B: 57, A: 255},
		color.RGBA{R: 253, G: 114, B: 57, A: 255},
		color.RGBA{R: 253, G: 114, B: 57, A: 255},
		color.RGBA{R: 253, G: 114, B: 57, A: 255},
		color.RGBA{R: 253, G: 114, B: 57, A: 255},
		color.RGBA{R: 253, G: 114, B: 57, A: 255},
		color.RGBA{R: 253, G: 114, B: 57, A: 255},
		color.RGBA{R: 253, G: 114, B: 57, A: 255},
		color.RGBA{R: 253, G: 114, B: 57, A: 255},
		color.RGBA{R: 253, G: 114, B: 57, A: 255},
		color.RGBA{R: 253, G: 114, B: 57, A: 255},
		color.RGBA{R: 253, G: 114, B: 57, A: 255},
		color.RGBA{R: 253, G: 114, B: 57, A: 255},
		color.RGBA{R: 253, G: 114, B: 57, A: 255},
		color.RGBA{R: 253, G: 114, B: 58, A: 255},
		color.RGBA{R: 253, G: 114, B: 58, A: 255},
		color.RGBA{R: 253, G: 114, B: 58, A: 255},
		color.RGBA{R: 254, G: 114, B: 59, A: 255},
		color.RGBA{R: 254, G: 114, B: 59, A: 255},
		color.RGBA{R: 254, G: 114, B: 59, A: 255},
		color.RGBA{R: 254, G: 114, B: 60, A: 255},
		color.RGBA{R: 254, G: 114, B: 60, A: 255},
		color.RGBA{R: 254, G: 114, B: 60, A: 255},
		color.RGBA{R: 251, G: 112, B: 59, A: 255},
		color.RGBA{R: 243, G: 108, B: 56, A: 255},
		color.RGBA{R: 236, G: 104, B: 52, A: 255},
		color.RGBA{R: 228, G: 100, B: 49, A: 255},
		color.RGBA{R: 221, G: 97, B: 46, A: 255},
		color.RGBA{R: 214, G: 93, B: 43, A: 255},
		color.RGBA{R: 206, G: 89, B: 40, A: 255},
		color.RGBA{R: 199, G: 85, B: 37, A: 255},
		color.RGBA{R: 191, G: 82, B: 33, A: 255},
		color.RGBA{R: 184, G: 78, B: 30, A: 255},
		color.RGBA{R: 176, G: 74, B: 27, A: 255},
		color.RGBA{R: 169, G: 70, B: 24, A: 255},
		color.RGBA{R: 162, G: 67, B: 21, A: 255},
		color.RGBA{R: 150, G: 62, B: 19, A: 255},
		color.RGBA{R: 139, G: 58, B: 18, A: 255},
		color.RGBA{R: 128, G: 54, B: 16, A: 255},
		color.RGBA{R: 117, G: 49, B: 15, A: 255},
		color.RGBA{R: 106, G: 45, B: 14, A: 255},
		color.RGBA{R: 95, G: 41, B: 12, A: 255},
		color.RGBA{R: 84, G: 36, B: 11, A: 255},
		color.RGBA{R: 73, G: 32, B: 10, A: 255},
		color.RGBA{R: 61, G: 28, B: 8, A: 255},
		color.RGBA{R: 50, G: 23, B: 7, A: 255},
		color.RGBA{R: 39, G: 19, B: 6, A: 255},
		color.RGBA{R: 28, G: 15, B: 4, A: 255},
		color.RGBA{R: 23, G: 13, B: 4, A: 255},
		color.RGBA{R: 23, G: 13, B: 4, A: 255},
		color.RGBA{R: 23, G: 13, B: 4, A: 255},
		color.RGBA{R: 23, G: 13, B: 4, A: 255},
		color.RGBA{R: 23, G: 13, B: 4, A: 255},
		color.RGBA{R: 23, G: 13, B: 4, A: 255},
		color.RGBA{R: 23, G: 13, B: 4, A: 255},
		color.RGBA{R: 23, G: 13, B: 4, A: 255},
		color.RGBA{R: 23, G: 13, B: 4, A: 255},
		color.RGBA{R: 23, G: 13, B: 4, A: 255},
		color.RGBA{R: 23, G: 13, B: 4, A: 255},
		color.RGBA{R: 23, G: 13, B: 4, A: 255},
	}*/
	WaterGradient = []color.Color{
		color.RGBA{R: 28, G: 146, B: 210, A: 255},
		color.RGBA{R: 30, G: 147, B: 210, A: 255},
		color.RGBA{R: 32, G: 148, B: 210, A: 255},
		color.RGBA{R: 34, G: 149, B: 211, A: 255},
		color.RGBA{R: 36, G: 150, B: 211, A: 255},
		color.RGBA{R: 38, G: 151, B: 212, A: 255},
		color.RGBA{R: 40, G: 152, B: 212, A: 255},
		color.RGBA{R: 42, G: 153, B: 213, A: 255},
		color.RGBA{R: 45, G: 154, B: 213, A: 255},
		color.RGBA{R: 47, G: 155, B: 213, A: 255},
		color.RGBA{R: 49, G: 156, B: 214, A: 255},
		color.RGBA{R: 51, G: 157, B: 214, A: 255},
		color.RGBA{R: 53, G: 158, B: 215, A: 255},
		color.RGBA{R: 55, G: 159, B: 215, A: 255},
		color.RGBA{R: 57, G: 160, B: 216, A: 255},
		color.RGBA{R: 60, G: 161, B: 216, A: 255},
		color.RGBA{R: 62, G: 162, B: 217, A: 255},
		color.RGBA{R: 64, G: 164, B: 217, A: 255},
		color.RGBA{R: 66, G: 165, B: 217, A: 255},
		color.RGBA{R: 68, G: 166, B: 218, A: 255},
		color.RGBA{R: 70, G: 167, B: 218, A: 255},
		color.RGBA{R: 72, G: 168, B: 219, A: 255},
		color.RGBA{R: 75, G: 169, B: 219, A: 255},
		color.RGBA{R: 77, G: 170, B: 220, A: 255},
		color.RGBA{R: 79, G: 171, B: 220, A: 255},
		color.RGBA{R: 81, G: 172, B: 221, A: 255},
		color.RGBA{R: 83, G: 173, B: 221, A: 255},
		color.RGBA{R: 85, G: 174, B: 221, A: 255},
		color.RGBA{R: 87, G: 175, B: 222, A: 255},
		color.RGBA{R: 90, G: 176, B: 222, A: 255},
		color.RGBA{R: 92, G: 177, B: 223, A: 255},
		color.RGBA{R: 94, G: 178, B: 223, A: 255},
		color.RGBA{R: 96, G: 179, B: 224, A: 255},
		color.RGBA{R: 98, G: 180, B: 224, A: 255},
		color.RGBA{R: 100, G: 182, B: 224, A: 255},
		color.RGBA{R: 102, G: 183, B: 225, A: 255},
		color.RGBA{R: 105, G: 184, B: 225, A: 255},
		color.RGBA{R: 107, G: 185, B: 226, A: 255},
		color.RGBA{R: 109, G: 186, B: 226, A: 255},
		color.RGBA{R: 111, G: 187, B: 227, A: 255},
		color.RGBA{R: 113, G: 188, B: 227, A: 255},
		color.RGBA{R: 115, G: 189, B: 228, A: 255},
		color.RGBA{R: 117, G: 190, B: 228, A: 255},
		color.RGBA{R: 120, G: 191, B: 228, A: 255},
		color.RGBA{R: 122, G: 192, B: 229, A: 255},
		color.RGBA{R: 124, G: 193, B: 229, A: 255},
		color.RGBA{R: 126, G: 194, B: 230, A: 255},
		color.RGBA{R: 128, G: 195, B: 230, A: 255},
		color.RGBA{R: 130, G: 196, B: 231, A: 255},
		color.RGBA{R: 132, G: 197, B: 231, A: 255},
		color.RGBA{R: 134, G: 199, B: 232, A: 255},
		color.RGBA{R: 137, G: 200, B: 232, A: 255},
		color.RGBA{R: 139, G: 201, B: 232, A: 255},
		color.RGBA{R: 141, G: 202, B: 233, A: 255},
		color.RGBA{R: 143, G: 203, B: 233, A: 255},
		color.RGBA{R: 145, G: 204, B: 234, A: 255},
		color.RGBA{R: 147, G: 205, B: 234, A: 255},
		color.RGBA{R: 149, G: 206, B: 235, A: 255},
		color.RGBA{R: 152, G: 207, B: 235, A: 255},
		color.RGBA{R: 154, G: 208, B: 235, A: 255},
		color.RGBA{R: 156, G: 209, B: 236, A: 255},
		color.RGBA{R: 158, G: 210, B: 236, A: 255},
		color.RGBA{R: 160, G: 211, B: 237, A: 255},
		color.RGBA{R: 162, G: 212, B: 237, A: 255},
		color.RGBA{R: 164, G: 213, B: 238, A: 255},
		color.RGBA{R: 167, G: 214, B: 238, A: 255},
		color.RGBA{R: 169, G: 215, B: 239, A: 255},
		color.RGBA{R: 171, G: 217, B: 239, A: 255},
		color.RGBA{R: 173, G: 218, B: 239, A: 255},
		color.RGBA{R: 175, G: 219, B: 240, A: 255},
		color.RGBA{R: 177, G: 220, B: 240, A: 255},
		color.RGBA{R: 179, G: 221, B: 241, A: 255},
		color.RGBA{R: 182, G: 222, B: 241, A: 255},
		color.RGBA{R: 184, G: 223, B: 242, A: 255},
		color.RGBA{R: 186, G: 224, B: 242, A: 255},
		color.RGBA{R: 188, G: 225, B: 243, A: 255},
		color.RGBA{R: 190, G: 226, B: 243, A: 255},
		color.RGBA{R: 192, G: 227, B: 243, A: 255},
		color.RGBA{R: 194, G: 228, B: 244, A: 255},
		color.RGBA{R: 197, G: 229, B: 244, A: 255},
		color.RGBA{R: 199, G: 230, B: 245, A: 255},
		color.RGBA{R: 201, G: 231, B: 245, A: 255},
		color.RGBA{R: 203, G: 232, B: 246, A: 255},
		color.RGBA{R: 205, G: 233, B: 246, A: 255},
		color.RGBA{R: 207, G: 235, B: 246, A: 255},
		color.RGBA{R: 209, G: 236, B: 247, A: 255},
		color.RGBA{R: 212, G: 237, B: 247, A: 255},
		color.RGBA{R: 214, G: 238, B: 248, A: 255},
		color.RGBA{R: 216, G: 239, B: 248, A: 255},
		color.RGBA{R: 218, G: 240, B: 249, A: 255},
		color.RGBA{R: 220, G: 241, B: 249, A: 255},
		color.RGBA{R: 222, G: 242, B: 250, A: 255},
		color.RGBA{R: 224, G: 243, B: 250, A: 255},
		color.RGBA{R: 227, G: 244, B: 250, A: 255},
		color.RGBA{R: 229, G: 245, B: 251, A: 255},
		color.RGBA{R: 231, G: 246, B: 251, A: 255},
		color.RGBA{R: 233, G: 247, B: 252, A: 255},
		color.RGBA{R: 235, G: 248, B: 252, A: 255},
		color.RGBA{R: 237, G: 249, B: 253, A: 255},
		color.RGBA{R: 239, G: 250, B: 253, A: 255},
	}
)

//go:embed data
var fs embed.FS

func Init() {
	// Slicing and loading wang blob tileset.

	/*road := load.Image(fs, "data/image/wall.png")
	blobW, blobH := 8, 6
	for y := 0; y < blobH; y++ {
		for x := 0; x < blobW; x++ {
			Images["wall_"+strconv.Itoa(x+y*blobW)] = split.Single(
				road, x*Tilesize, y*Tilesize,
				Tilesize, Tilesize, false,
			)
		}
	}*/

	/*
		// Slicing and loading player frames.
		player := load.Image(fs, "data/image/player.png")
		Images["player_idle"] = split.Single(player, 0, 0, PlayerW, PlayerH, false)
		for f := 0; f < 2; f++ {
			flipped := f == 1
			postfix := "right"
			if flipped {
				postfix = "left"
			}
			Images["player_look_"+postfix] = split.Single(player, PlayerW, 0, PlayerW, PlayerH, flipped)
			Images["player_jump_"+postfix] = split.Single(player, PlayerW*2, 0, PlayerW, PlayerH, flipped)
			Images["player_run_"+postfix] = split.Multi(player, PlayerW*3, 0, PlayerW, PlayerH, 4, flipped, Animation)

		}*/

	// Slicing and loading create image.
	//crate := load.Image(fs, "data/image/crate.png")
	//Images["crate"] = split.Single(crate, 0, 0, Cratesize, Cratesize, false)

	// Ship
	ship := load.Image(fs, "data/image/ship.png")
	Images["ship_idle"] = split.MultiRows(ship, 0, 0, ShipW, ShipH, 9, 40, false, Animation)

	carry := load.Image(fs, "data/image/ship_carry.png")
	Images["ship_carry"] = split.MultiRows(carry, 0, 0, ShipW, ShipH, 9, 40, false, Animation)

	ocean := load.Image(fs, "data/image/ocean.png")
	Images["ocean_idle"] = split.Single(ocean, 0, 0, OceanW, OceanH, false)

	rocks := load.Image(fs, "data/image/rocks.png")
	Images["rocks_idle"] = split.Single(rocks, 0, 0, RocksW, RocksH, false)

	island := load.Image(fs, "data/image/island.png")
	Images["island_idle"] = split.Single(island, 0, 0, RocksW, RocksH, false)
}
