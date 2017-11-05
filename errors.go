package wav

import "errors"

var (
	ErrFileFormat = errors.New("wav: invalid file format")

	ErrAudioFormat    = errors.New("wav: invalid AudioFormat")
	ErrChannels       = errors.New("wav: invalid Channels")
	ErrSampleRate     = errors.New("wav: invalid SampleRate")
	ErrBytesPerSample = errors.New("wav: invalid BytesPerSample")

	ErrFileReaderClosed = errors.New("wav: FileReader is closed or not created")
	ErrFileWriterClosed = errors.New("wav: FileWriter is closed or not created")
)
