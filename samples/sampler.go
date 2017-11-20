package samples

import (
	"errors"
	"io"
)

type (
	SampleReader interface {
		ReadSample(r io.Reader) (sample int32, err error)
	}

	SampleWriter interface {
		WriteSample(w io.Writer, sample int32) error
	}

	sampler interface {
		SampleReader
		SampleWriter
	}
)

func NewSampleReader(bitsPerSample int) (SampleReader, error) {
	return newSampler(bitsPerSample)
}

func NewSampleWriter(bitsPerSample int) (SampleWriter, error) {
	return newSampler(bitsPerSample)
}

func newSampler(bitsPerSample int) (sampler, error) {
	switch bitsPerSample {
	case 8:
		return new(sampler8), nil
	case 16:
		return new(sampler16), nil
	case 24:
		return new(sampler24), nil
	case 32:
		return new(sampler32), nil
	default:
		return nil, errors.New("invalid bitsPerSample")
	}
}

type sampler8 struct {
	buf [1]byte
}

var _ sampler = (*sampler8)(nil)

func (p *sampler8) ReadSample(r io.Reader) (sample int32, err error) {
	buf := p.buf[:]
	err = readAll(r, buf)
	if err != nil {
		return 0, err
	}
	sample = int32(le.GetInt8(buf))
	return
}

func (p *sampler8) WriteSample(w io.Writer, sample int32) error {
	buf := p.buf[:]
	i := int8(crop(sample, minInt8, maxInt8))
	le.PutInt8(buf, i)
	return writeAll(w, buf)
}

type sampler16 struct {
	buf [2]byte
}

var _ sampler = (*sampler16)(nil)

func (p *sampler16) ReadSample(r io.Reader) (sample int32, err error) {
	buf := p.buf[:]
	err = readAll(r, buf)
	if err != nil {
		return 0, err
	}
	sample = int32(le.GetInt16(buf))
	return
}

func (p *sampler16) WriteSample(w io.Writer, sample int32) error {
	buf := p.buf[:]
	i := int16(crop(sample, minInt16, maxInt16))
	le.PutInt16(buf, i)
	return writeAll(w, buf)
}

type sampler24 struct {
	buf [3]byte
}

var _ sampler = (*sampler24)(nil)

func (p *sampler24) ReadSample(r io.Reader) (sample int32, err error) {
	buf := p.buf[:]
	err = readAll(r, buf)
	if err != nil {
		return 0, err
	}
	sample = int32(le.GetInt24(buf))
	return
}

func (p *sampler24) WriteSample(w io.Writer, sample int32) error {
	buf := p.buf[:]
	i := int32(crop(sample, minInt24, maxInt24))
	le.PutInt24(buf, i)
	return writeAll(w, buf)
}

type sampler32 struct {
	buf [4]byte
}

var _ sampler = (*sampler32)(nil)

func (p *sampler32) ReadSample(r io.Reader) (sample int32, err error) {
	buf := p.buf[:]
	err = readAll(r, buf)
	if err != nil {
		return 0, err
	}
	sample = int32(le.GetInt32(buf))
	return
}

func (p *sampler32) WriteSample(w io.Writer, sample int32) error {
	buf := p.buf[:]
	i := int32(crop(sample, minInt32, maxInt32))
	le.PutInt32(buf, i)
	return writeAll(w, buf)
}
