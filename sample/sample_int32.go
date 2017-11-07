package sample

import (
	"errors"
	"io"
	"math"
)

type sampleReader interface {
	ReadSample(r io.Reader) (sample int32, err error)
}

type sampleWriter interface {
	WriteSample(w io.Writer, sample int32) error
}

type sampleReader8 struct {
	buf [1]byte
}

var _ sampleReader = (*sampleReader8)(nil)

func (sr *sampleReader8) ReadSample(r io.Reader) (sample int32, err error) {
	buf := sr.buf[:]
	err = readAll(r, buf)
	if err != nil {
		return 0, err
	}
	i := int8(buf[0])
	sample = int32(i)
	return
}

func (sr *sampleReader8) WriteSample(w io.Writer, sample int32) error {
	sample = crop(sample, math.MinInt8, math.MaxInt8)
	i := int8(sample)
	buf := sr.buf[:]
	buf[0] = byte(i)
	return writeAll(w, buf)
}

type sampleReader16 struct {
	buf [2]byte
}

var _ sampleReader = (*sampleReader16)(nil)

func (sr *sampleReader16) ReadSample(r io.Reader) (sample int32, err error) {
	buf := sr.buf[:]
	err = readAll(r, buf)
	if err != nil {
		return 0, err
	}
	sample = int32(le.GetInt16(buf))
	return
}

type sampleReader24 struct {
	buf [3]byte
}

var _ sampleReader = (*sampleReader24)(nil)

func (sr *sampleReader24) ReadSample(r io.Reader) (sample int32, err error) {
	buf := sr.buf[:]
	err = readAll(r, buf)
	if err != nil {
		return 0, err
	}
	sample = le.GetInt24(buf)
	return
}

type sampleReader32 struct {
	buf [4]byte
}

var _ sampleReader = (*sampleReader32)(nil)

func (sr *sampleReader32) ReadSample(r io.Reader) (sample int32, err error) {
	buf := sr.buf[:]
	err = readAll(r, buf)
	if err != nil {
		return 0, err
	}
	sample = le.GetInt32(buf)
	return
}

type SampleReader struct {
	r  io.Reader
	sr sampleReader
}

func NewSampleReader(r io.Reader, bitsPerSample int) (*SampleReader, error) {
	var sr sampleReader
	switch bitsPerSample {
	case 8:
		sr = new(sampleReader8)
	case 16:
		sr = new(sampleReader16)
	case 24:
		sr = new(sampleReader24)
	case 32:
		sr = new(sampleReader32)
	default:
		return nil, errors.New("wrong bitsPerSample")
	}
	return &SampleReader{
		r:  r,
		sr: sr,
	}, nil
}

func (sr *SampleReader) ReadSample() (sample int32, err error) {
	return sr.sr.ReadSample(sr.r)
}
