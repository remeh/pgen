package main

import (
	"log"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	// init sdl

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatalf("err:", err)
	}

	defer sdl.Quit()

	// create the surface we'll render on

	surface, err := sdl.CreateRGBSurfaceWithFormat(0, 1920, 1080, 32, uint32(sdl.PIXELFORMAT_RGBA32))
	if err != nil {
		log.Fatalf("err: can't create the dest surface: %v", err)
	}
	defer surface.Free()

	// start drawing on the surface

	if err := render(surface); err != nil {
		log.Fatalf("err: %v", err)
	}

	// store the result in a file

	img.SavePNG(surface, "out.png")
}
