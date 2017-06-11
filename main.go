package main

import (
	"flag"
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var printOctave = flag.Bool("octaves", false, "also outputs the first three octaves (creates 4x bigger outputs)")
var zoomFlag = flag.Float64("zoom", 8, "zoom on the terrain")
var xFlag = flag.Int("x", 1280, "the horizontal size")
var yFlag = flag.Int("y", 960, "the vertical size")
var seedFlag = flag.Int64("seed", 0, "the seed for the noise")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		_ = pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	zoom := *zoomFlag
	x := *xFlag
	y := *yFlag
	seed := *seedFlag

	var tot *image.RGBA
	t := NewTerrain(3, zoom, 0.4, seed)

	if *printOctave {
		m := t.Render(x, y)
		o := t.RenderOctaves(x, y)

		tot = image.NewRGBA(image.Rect(0, 0, x*2+1, y*2+1))

		draw.Draw(tot, image.Rect(0, 0, x, y), m, image.ZP, draw.Src)
		draw.Draw(tot, image.Rect(x+1, 0, x*2+1, y), o[0], image.ZP, draw.Src)
		draw.Draw(tot, image.Rect(0, y+1, x, y*2+1), o[1], image.ZP, draw.Src)
		draw.Draw(tot, image.Rect(x+1, y+1, x*2+1, y*2+1), o[2], image.ZP, draw.Src)
	} else {
		tot = t.Render(x, y)
	}
	outfile, err := os.OpenFile("out.png", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("failed to output to file: " + err.Error())
	}
	err = png.Encode(outfile, tot)
	if err != nil {
		log.Fatal("failed to output to file: " + err.Error())
	}
	/*
		for i, img := range t.RenderOctaves(x, y) {
			outfile, err := os.OpenFile("out_"+strconv.Itoa(i+1)+".png", os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal("failed to output to file: " + err.Error())
			}
			err = png.Encode(outfile, img)
			if err != nil {
				log.Fatal("failed to output to file: " + err.Error())
			}

		}
	*/
}
