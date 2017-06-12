package main

import (
	"image"
	"image/color"
	"image/draw"
	"log"
	"runtime"

	colorful "github.com/lucasb-eyer/go-colorful"
)

type Terrain2D struct {
	octaves []*OctaveNoise

	//Sealevel varies between 0 and 1 and determines where to draw blue
	sealevel float64

	scalefactor float64
}

func NewTerrain(octavesAmount int, zoom, sealevel float64, seed int64) *Terrain2D {
	var octaves []*OctaveNoise
	for i := 1; i <= octavesAmount; i += 1 {
		//octaves = append(octaves, NewOctaveNoise(math.Pow(2, float64(i))*scalefactor, seed))
		octaves = append(octaves, NewOctaveNoise(float64(i*i*2), 0.02/zoom, seed))
	}

	var scale float64
	for i := 0; i < octavesAmount; i++ {
		scale += 1 / octaves[i].octaveindex
	}

	return &Terrain2D{
		octaves:     octaves,
		sealevel:    sealevel,
		scalefactor: scale,
	}
}

//Returns the terrain height value normalized between -1 and 1
func (t *Terrain2D) Height(x, y float64) float64 {
	var sum float64
	for _, octave := range t.octaves {
		sum += octave.Eval2(x, y)
	}
	sum /= t.scalefactor
	return sum
}

func (t *Terrain2D) Render(x, y int) *image.RGBA {
	return t.RenderZoom(x, y, 1)
}

func (t *Terrain2D) RenderZoom(x, y int, zoom float64) *image.RGBA {
	m := image.NewRGBA(image.Rect(0, 0, x, y))
	slices := make(chan *image.RGBA)
	amount := runtime.GOMAXPROCS(0)
	for i := 0; i < amount; i++ {
		go func(i int) {
			r := image.Rect((x/amount)*i, 0, (x/amount)*(i+1), y)
			m := image.NewRGBA(r)
			for i := r.Min.X; i < r.Max.X; i++ {
				for j := 0; j < y; j++ {
					// height is divided by 2 to have a range of 1.0 instead of 2.0 (-1,+1)
					height := t.Height(float64(i), float64(j))/2 + 0.5
					m.Set(i, j, colorizer(height, t.sealevel))
				}
			}
			slices <- m
		}(i)
	}
	for i := 0; i < amount; i++ {
		tmp := <-slices
		log.Printf("Composing slice with x: %d → %d", tmp.Rect.Min.X, tmp.Rect.Max.X)
		draw.Draw(m, tmp.Rect, tmp, tmp.Rect.Min, draw.Src)
	}
	return m
}

func (t *Terrain2D) RenderOctaves(x, y int) (images []*image.RGBA) {
	for i := 0; i < len(t.octaves); i++ {
		images = append(images, image.NewRGBA(image.Rect(0, 0, x, y)))
	}
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			for o, octave := range t.octaves {
				// height is divided by 2 to have a range of 1.0 instead of 2.0 (-1,+1)
				height := octave.Eval2(float64(i), float64(j))*octave.octaveindex/2 + 0.5
				images[o].Set(i, j, colorizer(height, 0))
			}
		}
	}
	return
}

//value must be between 0 and 1
func colorizer(value, sealevel float64) *color.RGBA {
	if value < sealevel {
		return &color.RGBA{0, 70, 200, 255}
	}
	r, g, b := colorful.Hsl(280+-value*360, 0.5, 0.5).RGB255()
	return &color.RGBA{r, g, b, 255}
}
