package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

var textures []*sdl.Surface

type Renderer interface {
	Render(surface *sdl.Surface) error
}

func loadRes() error {
	// load every PNGs and JPGs available
	// in res/textures
	files, err := ioutil.ReadDir("res/textures")
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if strings.HasSuffix(file.Name(), "png") ||
			strings.HasSuffix(file.Name(), "jpg") {
			s, err := img.Load("res/textures/" + file.Name())
			if err != nil {
				log.Printf("warn: error while loading %s: %v\n", file.Name(), err)
				continue
			}
			textures = append(textures, s)
		}
	}

	if len(textures) == 0 {
		return fmt.Errorf("didn't find any textures in res/textures/")
	}

	log.Println("loaded", len(textures), "textures")

	return nil
}
