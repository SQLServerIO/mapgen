package main

import (
	"image/png"
	"log"
	"os"
)

func main() {
	var seed int64
	var scalefactor float64

	scalefactor = 0.005
	x := 640
	y := 480

	t := NewTerrain(4, scalefactor, 0.4, seed)
	m := t.Render(x, y)
	outfile, err := os.OpenFile("out.png", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("failed to output to file: " + err.Error())
	}
	err = png.Encode(outfile, m)
	if err != nil {
		log.Fatal("failed to output to file: " + err.Error())
	}
}
