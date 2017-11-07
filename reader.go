package wav

import (
	"encoding/binary"
	"errors"
	"io"
	"os"
)

type FileReader struct {
	config     Config
	file       *os.File
	dataLength int
}

func NewFileReader(fileName string) (*FileReader, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	fr := &FileReader{file: file}
	if err = fr.readChunks(); err != nil {
		file.Close()
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
	fr.dataLength = 0
	return err
}

func (fr *FileReader) Read(data []byte) (n int, err error) {
	if fr.dataLength == 0 {
		return 0, nil
	}
	n = len(data)
	if n > fr.dataLength {
		n = fr.dataLength
	}
	n, err = fr.file.Read(data[:n])
	fr.dataLength -= n
	return n, err
}

func (fr *FileReader) Config() Config {
	return fr.config
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
	var ch chunkHeader

	//-----------------------------------------------------
	// RIFF header
	n, err := read(fr.file, binary.LittleEndian, &ch)
	if err != nil {
		return err
	}
	pos += int64(n)
	if ch.Id != tag_RIFF {
		return ErrFileFormat
	}
	fileSize := int64(n) + int64(ch.Size)

	//-----------------------------------------------------
	// format WAVE
	var format tag
	n, err = read(fr.file, binary.LittleEndian, &format)
	if err != nil {
		return err
	}
	pos += int64(n)
	if format != tag_WAVE {
		return errors.New("wav: format is not WAVE")
	}

	//-----------------------------------------------------
	// read all chunks

	var chunks = make(map[tag]*chunkLoc)

	for {
		n, err = read(fr.file, binary.LittleEndian, &ch)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		pos += int64(n)

		loc := chunkLoc{
			pos:  pos,
			size: int64(ch.Size),
		}

		_, err = fr.file.Seek(loc.size, os.SEEK_CUR)
		if err != nil {
			return err
		}
		pos += loc.size // chunk data

		chunks[ch.Id] = &loc
	}

	// check file size
	if pos != fileSize {
		return ErrFileFormat
	}

	//-----------------------------------------------------
	// chunk fmt_
	loc, ok := chunks[tag_fmt_]
	if !ok {
		return errors.New("wav: has not chunk \"fmt \"")
	}
	_, err = fr.file.Seek(loc.pos, os.SEEK_SET)
	if err != nil {
		return err
	}
	var fmt_data fmtData
	err = binary.Read(fr.file, binary.LittleEndian, &fmt_data)
	if err != nil {
		return err
	}
	fr.config = fmtDataToConfig(fmt_data)

	//-----------------------------------------------------
	// chunk data
	loc, ok = chunks[tag_data]
	if !ok {
		return errors.New("wav: has not chunk \"data\"")
	}
	_, err = fr.file.Seek(loc.pos, os.SEEK_SET)
	if err != nil {
		return err
	}
	fr.dataLength = int(loc.size)

	return nil
}

func read(r io.Reader, order binary.ByteOrder, data interface{}) (n int, err error) {
	cr := &countReader{r: r}
	err = binary.Read(cr, order, data)
	n = cr.n
	return
}

type countReader struct {
	r io.Reader
	n int
}

func (p *countReader) Read(data []byte) (n int, err error) {
	n, err = p.r.Read(data)
	p.n += n
	return
}
