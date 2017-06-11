package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func main() {
	var seed int64
	var scalefactor float64

	scalefactor = 0.01
	x := 640
	y := 480

	var octaves []*OctaveNoise
	for i := 1; i < 3; i += 1 {
		octaves = append(octaves, NewOctaveNoise(math.Pow(2, float64(i))*scalefactor, seed))
	}

	m := image.NewRGBA(image.Rect(0, 0, x, y))
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			var sum float64
			for _, octave := range octaves {
				sum += octave.Eval2(float64(i), float64(j))
			}
			sum /= float64(len(octaves))
			isum := uint8(float64(255) * sum)
			m.Set(i, j, color.RGBA{isum, isum, isum, 255})
		}
	}
	outfile, err := os.OpenFile("out.png", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("failed to output to file: " + err.Error())
	}
	err = png.Encode(outfile, m)
	if err != nil {
		log.Fatal("failed to output to file: " + err.Error())
	}
}
