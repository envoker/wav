package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/envoker/wav"
)

func main() {
	const fileName = "test.wav"
	TestWaveWrite(fileName)
	TestWaveRead(fileName)
	TextFileHexDump(fileName)
}

func TestWaveWrite(fileName string) {

	c := wav.Config{
		AudioFormat:   wav.WAVE_FORMAT_PCM,
		Channels:      2,
		SampleRate:    22050,
		BitsPerSample: 16,
	}

	fw, err := wav.NewFileWriter(fileName, c)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fw.Close()

	d := make([]byte, 2048)

	for i := range d {
		d[i] = byte(i)
	}

	fw.Write(d)
}

func TestWaveRead(fileName string) {

	fr, err := wav.NewFileReader(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fr.Close()

	c, _ := fr.Config()
	fmt.Printf("%+v\n", c)
}

func TextFileHexDump(fileName string) error {

	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	data := make([]byte, 100)

	n, err := f.Read(data)
	if err != nil {
		return err
	}

	fmt.Println(hex.Dump(data[:n]))

	return nil
}
