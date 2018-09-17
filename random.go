package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type Random struct {
	particles int
}

func NewRandomRenderer(particles int) *Random {
	rand.Seed(time.Now().Unix())
	return &Random{
		particles: particles,
	}
}

func (r *Random) Render(surface *sdl.Surface) error {
	if err := loadRes(); err != nil {
		return fmt.Errorf("while loading resources: %v", err)
	}

	// fill the background with dark
	rect := sdl.Rect{0, 0, surface.W, surface.H}
	surface.FillRect(&rect, 0x151515ff)

	for i := 0; i < r.particles; i++ {
		if err := r.one(surface); err != nil {
			return err
		}
	}

	return nil
}

func (r *Random) one(surface *sdl.Surface) error {

	p := rand.Intn(len(textures))
	t := textures[p]

	var w, h, s int32

	for {
		w = int32(rand.Intn(int(surface.W)))
		h = int32(rand.Intn(int(surface.H)))
		s = 16 + int32(rand.Intn(72))

		if w+s < surface.W && h+s < surface.H {
			break
		}
	}

	t.SetAlphaMod(uint8(rand.Int()))

	return t.BlitScaled(nil, surface, &sdl.Rect{w, h, s, s})
}
