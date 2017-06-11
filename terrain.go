package main

import (
	"image"
	"image/color"

	colorful "github.com/lucasb-eyer/go-colorful"
)

type Terrain2D struct {
	octaves []*OctaveNoise

	//Sealevel varies between 0 and 1 and determines where to draw blue
	sealevel float64
}

func NewTerrain(octavesAmount int, scalefactor, sealevel float64, seed int64) *Terrain2D {
	var octaves []*OctaveNoise
	for i := 1; i <= octavesAmount; i += 1 {
		//octaves = append(octaves, NewOctaveNoise(math.Pow(2, float64(i))*scalefactor, seed))
		octaves = append(octaves, NewOctaveNoise(float64(i)*scalefactor, seed))
	}
	return &Terrain2D{
		octaves:  octaves,
		sealevel: sealevel,
	}
}

//Returns the terrain height value normalized between -1 and 1
func (t *Terrain2D) Height(x, y float64) float64 {
	var sum float64
	for _, octave := range t.octaves {
		sum += octave.Eval2(x, y)
	}
	sum /= float64(len(t.octaves))
	return sum
}

func (t *Terrain2D) Render(x, y int) *image.RGBA {
	return t.RenderZoom(x, y, 1)
}

func (t *Terrain2D) RenderZoom(x, y int, zoom float64) *image.RGBA {
	m := image.NewRGBA(image.Rect(0, 0, x, y))
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			// height is divided by 2 to have a range of 1.0 instead of 2.0 (-1,+1)
			height := t.Height(float64(i), float64(j))/2 + 0.5
			m.Set(i, j, colorizer(height, t.sealevel))
		}
	}
	return m
}

//value must be between 0 and 1
func colorizer(value, sealevel float64) *color.RGBA {
	if value < sealevel {
		return &color.RGBA{0, 70, 200, 255}
	}
	r, g, b := colorful.Hsl(280+-value*360, 0.5, 0.5).RGB255()
	return &color.RGBA{r, g, b, 255}
}
