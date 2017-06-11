package main

import noise "github.com/ojrac/opensimplex-go"

type OctaveNoise struct {
	*noise.Noise
	octave float64
}

func NewOctaveNoise(octave float64, seed int64) *OctaveNoise {
	return &OctaveNoise{
		Noise:  noise.NewWithSeed(seed),
		octave: octave,
	}
}

func (on *OctaveNoise) Eval2(x, y float64) float64 {
	return on.Noise.Eval2(
		x*on.octave,
		y*on.octave,
	)
}
func (on *OctaveNoise) Eval3(x, y, z float64) float64 {
	return on.Noise.Eval3(
		x*on.octave,
		y*on.octave,
		z*on.octave,
	)
}
func (on *OctaveNoise) Eval4(x, y, z, w float64) float64 {
	return on.Noise.Eval4(
		x*on.octave,
		y*on.octave,
		z*on.octave,
		w*on.octave,
	)
}
