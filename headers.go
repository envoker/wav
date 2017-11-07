package wav

type tag [4]byte

var (
	tag_RIFF = tag{'R', 'I', 'F', 'F'} // "RIFF"
	tag_WAVE = tag{'W', 'A', 'V', 'E'} // "WAVE"
	tag_fmt_ = tag{'f', 'm', 't', ' '} // "fmt "
	tag_data = tag{'d', 'a', 't', 'a'} // "data"
)

type chunkHeader struct {
	Id   tag
	Size uint32
}

type fmtData struct {
	AudioFormat   uint16
	Channels      uint16
	SampleRate    uint32
	BytesPerSec   uint32
	BytesPerBlock uint16
	BitsPerSample uint16
}
