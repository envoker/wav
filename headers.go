package wav

import (
	"encoding/binary"
)

const (
	sizeRiffFormat = 4 // WAVE
)

var (
	sizeRiffHeader  = binary.Size(riffHeader{})
	sizeChunkHeader = binary.Size(chunkHeader{})
	sizeFmtData     = binary.Size(fmtData{})
)

type tag [4]byte

//var (
//	tag_RIFF = strToTag("RIFF")
//	tag_WAVE = strToTag("WAVE")
//	tag_fmt_ = strToTag("fmt ")
//	tag_data = strToTag("data")
//)

//func strToTag(s string) (t tag) {
//	copy(t[:], []byte(s))
//	return
//}

var (
	tag_RIFF = tag{'R', 'I', 'F', 'F'}
	tag_WAVE = tag{'W', 'A', 'V', 'E'}
	tag_fmt_ = tag{'f', 'm', 't', ' '}
	tag_data = tag{'d', 'a', 't', 'a'}
)

type riffHeader struct {
	Id     tag
	Size   uint32
	Format tag
}

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
