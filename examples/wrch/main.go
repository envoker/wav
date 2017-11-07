package main

import (
	"fmt"
	"log"

	"github.com/envoker/wav"
)

func main() {
	fileName := "./test.wav"

	var c1 = wav.Config{
		AudioFormat:   wav.WAVE_FORMAT_PCM,
		Channels:      1,
		SampleRate:    8000,
		BitsPerSample: 16,
	}
	fw, err := wav.NewFileWriter(fileName, c1)
	if err != nil {
		log.Fatal(err)
	}
	fw.Write([]byte{0xf5, 0x11, 0x7b})
	fw.Close()

	fmt.Println(c1)

	fr, err := wav.NewFileReader(fileName)
	if err != nil {
		log.Fatal(err)
	}
	c2 := fr.Config()

	var data [64]byte
	n, err := fr.Read(data[:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read data: % x\n", data[:n])
	fr.Close()

	fmt.Println(c2)
}
