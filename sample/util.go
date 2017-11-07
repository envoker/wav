package sample

import (
	"errors"
	"io"
)

func readAll(r io.Reader, data []byte) error {
	n, err := r.Read(data)
	if err != nil {
		return err
	}
	if n < len(data) {
		return errors.New("insufficient data length")
	}
	return nil
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
