package wav

const (
	sizeofUint16 = 2
	sizeofUint32 = 4

	sizeofChunkId    = sizeofUint32
	sizeofChunkSize  = sizeofUint32
	sizeofRiffFormat = sizeofUint32

	sizeofRiffHeader  = sizeofChunkId + sizeofChunkSize + sizeofRiffFormat
	sizeofChunkHeader = sizeofChunkId + sizeofChunkSize

	sizeofFmtData = sizeofUint16 + // AudioFormat
		sizeofUint16 + // Channels
		sizeofUint32 + // SampleRate
		sizeofUint32 + // BytesPerSec
		sizeofUint16 + // BytesPerBlock
		sizeofUint16 // BitsPerSample
)

type token [4]byte

var (
	token_RIFF = strToken("RIFF")
	token_WAVE = strToken("WAVE")
	token_fmt  = strToken("fmt ")
	token_data = strToken("data")
)

func strToken(s string) (t token) {
	copy(t[:], []byte(s))
	return
}

type riffHeader struct {
	Id     token
	Size   uint32
	Format token
}

type chunkHeader struct {
	Id   token
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
