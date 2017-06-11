package main

import noise "github.com/ojrac/opensimplex-go"

type OctaveNoise struct {
	*noise.Noise
	octaveindex float64
	octave      float64
}

func NewOctaveNoise(octave, scalefactor float64, seed int64) *OctaveNoise {
	return &OctaveNoise{
		Noise:       noise.NewWithSeed(seed),
		octaveindex: octave,
		octave:      octave * scalefactor,
	}
}

func (on *OctaveNoise) Eval2(x, y float64) float64 {
	return on.Noise.Eval2(
		x*on.octave,
		y*on.octave,
	) / on.octaveindex
}

func (on *OctaveNoise) Eval3(x, y, z float64) float64 {
	return on.Noise.Eval3(
		x*on.octave,
		y*on.octave,
		z*on.octave,
	) / on.octaveindex
}

func (on *OctaveNoise) Eval4(x, y, z, w float64) float64 {
	return on.Noise.Eval4(
		x*on.octave,
		y*on.octave,
		z*on.octave,
		w*on.octave,
	) / on.octaveindex
}
