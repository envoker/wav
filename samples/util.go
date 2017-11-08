package samples

import (
	"io"
)

func readAll(r io.Reader, data []byte) error {
	_, err := io.ReadFull(r, data)
	return err
}

func writeAll(w io.Writer, data []byte) error {
	_, err := w.Write(data)
	return err
}

func crop(x int32, min, max int32) int32 {
	if x < min {
		x = min
	}
	if x > max {
		x = max
	}
	return x
}
