package wav

type Config struct {
	AudioFormat    int // тип формата (1 - PCM; 6 - A-law, 7 - Mu-law)
	Channels       int // количество каналов (1 - моно; 2 - стeрео)
	SampleRate     int // частота дискретизации (8000, ...)
	BytesPerSample int // 1, 2, 3, 4
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

	switch c.BytesPerSample {
	case 1, 2, 3, 4:
	default:
		return ErrBytesPerSample
	}

	return nil
}

func (c *Config) BytesPerSec() int {
	return c.Channels * c.BytesPerSample * c.SampleRate
}

func (c *Config) BytesPerBlock() int {
	return c.Channels * c.BytesPerSample
}

func configToFmtData(c Config) fmtData {
	return fmtData{
		AudioFormat:   uint16(c.AudioFormat),
		Channels:      uint16(c.Channels),
		SampleRate:    uint32(c.SampleRate),
		BitsPerSample: uint16(c.BytesPerSample * 8),
		BytesPerSec:   uint32(c.BytesPerSec()),
		BytesPerBlock: uint16(c.BytesPerBlock()),
	}
}

func fmtDataToConfig(d fmtData) Config {
	return Config{
		AudioFormat:    int(d.AudioFormat),
		Channels:       int(d.Channels),
		SampleRate:     int(d.SampleRate),
		BytesPerSample: int(d.BitsPerSample) / 8,
	}
}
