package main

import (
	"bufio"
	"time"

	"github.com/envoker/wav"
	"github.com/envoker/wav/samples"
)

func GenerateWave(fileName string, duration time.Duration, sampleRate float32, bitsPerSample int, samplers []NextSampler) error {

	Tmax := float32(duration.Seconds())

	c := wav.Config{
		AudioFormat:   wav.WAVE_FORMAT_PCM,
		Channels:      len(samplers),
		SampleRate:    int(sampleRate),
		BitsPerSample: bitsPerSample,
	}

	fw, err := wav.NewFileWriter(fileName, c)
	if err != nil {
		return err
	}
	defer fw.Close()

	bw := bufio.NewWriterSize(fw, c.BytesPerSec())

	sw, err := samples.NewSampleWriter(c.BitsPerSample)
	if err != nil {
		return err
	}
	maxValue := samples.MaxSample(c.BitsPerSample)
	maxValueFloat := float32(maxValue)
	n := int(Tmax * sampleRate)
	for i := 0; i < n; i++ {
		for _, sampler := range samplers {
			sample := int32(sampler.NextSample() * maxValueFloat)
			if err = sw.WriteSample(fw, sample); err != nil {
				return err
			}
		}
	}

	return bw.Flush()
}
