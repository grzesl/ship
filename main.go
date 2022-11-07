package main

import (
	"log"

	"github.com/grzesl/ship/scene"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"
)

func main() {

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowSizeLimits(300, 200, -1, -1)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOn)
	ebiten.SetWindowResizable(true)

	if err := ebiten.RunGame(engine.NewGame(&scene.Game{})); err != nil {
		log.Fatal(err)
	}
}
