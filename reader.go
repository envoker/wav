package wav

import (
	"encoding/binary"
	"io"
	"os"
)

type FileReader struct {
	config Config
	file   *os.File
	r      io.Reader
}

func NewFileReader(fileName string) (*FileReader, error) {

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	fr := &FileReader{
		file: file,
	}

	if err = fr.readChunks(); err != nil {
		return nil, err
	}

	return fr, nil
}

func (fr *FileReader) Close() error {

	if fr.file == nil {
		return ErrFileReaderClosed
	}

	err := fr.file.Close()
	fr.file = nil
	return err
}

func (fr *FileReader) Read(data []byte) (n int, err error) {
	return fr.r.Read(data)
}

func (fr *FileReader) Config() (Config, error) {
	return fr.config, nil
}

type chunkLoc struct {
	pos  int64
	size int64
}

func (fr *FileReader) readChunks() error {

	_, err := fr.file.Seek(0, os.SEEK_SET)
	if err != nil {
		return err
	}

	var pos int64
	var rh riffHeader

	err = binary.Read(fr.file, binary.LittleEndian, &rh)
	if err != nil {
		return err
	}
	pos += int64(sizeRiffHeader)

	if rh.Id != tag_RIFF {
		return ErrFileFormat
	}
	if rh.Format != tag_WAVE {
		return ErrFileFormat
	}
	fileSize := int64(sizeRiffHeader-sizeRiffFormat) + int64(rh.Size)

	var chunks = make(map[tag]*chunkLoc)
	var ch chunkHeader

	for {
		err = binary.Read(fr.file, binary.LittleEndian, &ch)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		pos += int64(sizeChunkHeader)

		cl := chunkLoc{
			pos:  pos,
			size: int64(ch.Size),
		}

		_, err = fr.file.Seek(cl.size, os.SEEK_CUR)
		if err != nil {
			return err
		}
		pos += cl.size // chunk data

		chunks[ch.Id] = &cl
	}

	if pos != fileSize {
		return ErrFileFormat
	}

	// fmt_
	chunkFmt, ok := chunks[tag_fmt_]
	if !ok {
		return ErrFileFormat
	}
	err = fr.readChunkFmt(chunkFmt)
	if err != nil {
		return err
	}

	// data
	chunkData, ok := chunks[tag_data]
	if !ok {
		return ErrFileFormat
	}
	err = fr.readChunkData(chunkData)
	if err != nil {
		return err
	}

	return nil
}

func (fr *FileReader) readChunkFmt(cl *chunkLoc) error {
	_, err := fr.file.Seek(cl.pos, os.SEEK_SET)
	if err != nil {
		return err
	}
	var c_data fmtData
	err = binary.Read(fr.file, binary.LittleEndian, &c_data)
	if err != nil {
		return err
	}
	fr.config = fmtDataToConfig(c_data)
	return nil
}

func (fr *FileReader) readChunkData(cl *chunkLoc) error {
	_, err := fr.file.Seek(cl.pos, os.SEEK_SET)
	if err != nil {
		return err
	}
	fr.r = newLimitReader(fr.file, int(cl.size))
	//fr.r = bufio.NewReaderSize(fr.file, int(cl.size))
	return nil
}

type limitReader struct {
	r          io.Reader
	dataLength int
}

func newLimitReader(r io.Reader, n int) *limitReader {
	return &limitReader{
		r:          r,
		dataLength: n,
	}
}

func (lr *limitReader) Read(data []byte) (n int, err error) {

	if lr.dataLength == 0 {
		return 0, nil
	}

	n = len(data)
	if n > lr.dataLength {
		n = lr.dataLength
	}

	n, err = lr.r.Read(data[:n])
	lr.dataLength -= n

	return n, err
}
