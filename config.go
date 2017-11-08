package wav

type Config struct {
	AudioFormat   int // тип формата (1 - PCM; 6 - A-law, 7 - Mu-law)
	Channels      int // количество каналов (1 - моно; 2 - стeрео)
	SampleRate    int // частота дискретизации (8000, ...)
	BitsPerSample int // [8, 16, 24, 32]
}

func (c *Config) checkError() error {

	if (c.AudioFormat < 0) || (c.AudioFormat > 65535) {
		return ErrAudioFormat
	}

	if (c.Channels < 1) || (c.Channels > 32) {
		return ErrChannels
	}

	if (c.SampleRate < 10) || (c.SampleRate > 200000) {
		return ErrSampleRate
	}

	switch c.BitsPerSample {
	case 8, 16, 24, 32:
	default:
		return ErrBitsPerSample
	}

	return nil
}

func (c *Config) BytesPerSec() int {
	return c.SampleRate * c.Channels * c.BitsPerSample / 8
}

func (c *Config) BytesPerBlock() int {
	return c.Channels * c.BitsPerSample / 8
}

func configToFmtData(c Config) fmtData {
	return fmtData{
		AudioFormat:   uint16(c.AudioFormat),
		Channels:      uint16(c.Channels),
		SampleRate:    uint32(c.SampleRate),
		BytesPerSec:   uint32(c.BytesPerSec()),
		BytesPerBlock: uint16(c.BytesPerBlock()),
		BitsPerSample: uint16(c.BitsPerSample),
	}
}

func fmtDataToConfig(d fmtData) Config {
	return Config{
		AudioFormat:   int(d.AudioFormat),
		Channels:      int(d.Channels),
		SampleRate:    int(d.SampleRate),
		BitsPerSample: int(d.BitsPerSample),
	}
}
