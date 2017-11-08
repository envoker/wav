package main

import (
	"math"
)

type NextSampler interface {
	NextSample() float32 // [-1 ... +1]
}

type Garmonica struct {
	Amplitude, Frequency, Phase float32
}

type toneSampler struct {
	amplitude float32
	phase     float32
	w         float32
	t, dt     float32
}

func (ts *toneSampler) NextSample() float32 {

	u := float64(ts.w*ts.t + ts.phase)
	sample := ts.amplitude * float32(math.Sin(u))

	ts.t += ts.dt

	return sample
}

func NewToneSampler(g Garmonica, sampleRate float32) NextSampler {
	return &toneSampler{
		amplitude: g.Amplitude,
		phase:     g.Phase,
		w:         2 * math.Pi * g.Frequency,
		t:         0,
		dt:        1.0 / sampleRate,
	}
}

func MakeSamplers(gs []Garmonica, sampleRate float32) []NextSampler {
	samplers := make([]NextSampler, len(gs))
	for i, g := range gs {
		samplers[i] = NewToneSampler(g, sampleRate)
	}
	return samplers
}
