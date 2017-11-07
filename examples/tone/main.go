package main

import (
	"log"
	"math"
	"time"
)

const bitsPerSample = 24

func monoTone() error {

	sampleRate := float32(22050)

	gs := []Garmonica{
		Garmonica{
			Amplitude: 0.9,
			Frequency: 1000,
			Phase:     math.Pi * 0.5,
		},
	}

	samplers := MakeSamplers(gs, sampleRate)
	return GenerateWave("./mono-tone.wav", time.Second*9, sampleRate, bitsPerSample, samplers)
}

func stereoTone() error {

	sampleRate := float32(44100)

	gs := []Garmonica{
		Garmonica{
			Amplitude: 0.75,
			Frequency: 500,
			Phase:     math.Pi * 0.5,
		},
		Garmonica{
			Amplitude: 0.60,
			Frequency: 1000,
			Phase:     0,
		},
	}

	samplers := MakeSamplers(gs, sampleRate)
	return GenerateWave("./stereo-tone.wav", time.Second*13, sampleRate, bitsPerSample, samplers)
}

func multiTone() error {

	sampleRate := float32(44100)

	gs := []Garmonica{
		Garmonica{
			Amplitude: 0.75,
			Frequency: 1500,
			Phase:     math.Pi * 0.5,
		},
		Garmonica{
			Amplitude: 0.60,
			Frequency: 3000,
			Phase:     0,
		},
		Garmonica{
			Amplitude: 0.90,
			Frequency: 400,
			Phase:     0,
		},
		Garmonica{
			Amplitude: 0.50,
			Frequency: 1000,
			Phase:     0,
		},
	}

	samplers := MakeSamplers(gs, sampleRate)
	return GenerateWave("./multi-tone.wav", time.Second*7, sampleRate, bitsPerSample, samplers)
}

func main() {
	if err := monoTone(); err != nil {
		log.Fatal(err)
	}

	if err := stereoTone(); err != nil {
		log.Fatal(err)
	}

	if err := multiTone(); err != nil {
		log.Fatal(err)
	}
}
