package wav

import (
	"bufio"
	"encoding/binary"
	"io"
	"os"
)

type FileReader struct {
	config     Config
	dataLength int64
	file       *os.File
	r          io.Reader
}

func NewFileReader(fileName string, config *Config) (*FileReader, error) {

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	fr := &FileReader{
		dataLength: 0,
		file:       file,
	}

	if err = fr.readChunks(); err != nil {
		return nil, err
	}

	*config = fr.config

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

func (fr *FileReader) prevRead(data []byte) (n int, err error) {

	if fr.dataLength == 0 {
		return 0, nil
	}

	n = len(data)
	if n > int(fr.dataLength) {
		n = int(fr.dataLength)
	}

	n, err = fr.file.Read(data[:n])
	if err != nil {
		return 0, err
	}

	fr.dataLength -= int64(n)

	return
}

func (fr *FileReader) Read(data []byte) (n int, err error) {
	return fr.r.Read(data)
}

func (fr *FileReader) getConfig(c *Config) {
	*c = fr.config
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
	var riff_h riffHeader

	err = binary.Read(fr.file, binary.LittleEndian, &riff_h)
	if err != nil {
		return err
	}
	pos += sizeofRiffHeader

	if riff_h.Id != token_RIFF {
		return ErrFileFormat
	}
	if riff_h.Format != token_WAVE {
		return ErrFileFormat
	}
	fileSize := int64(riff_h.Size) + sizeofChunkHeader

	//fmt.Printf("riff id: <%s>\n", riff_h.Id)
	//fmt.Printf("format: <%s>\n", riff_h.Format)
	//fmt.Printf("chunk id: <%s>\n", token_WAVE)

	var chunks = make(map[token]*chunkLoc)
	var ch chunkHeader

	for {
		err = binary.Read(fr.file, binary.LittleEndian, &ch)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		pos += sizeofChunkHeader // chunk header

		c := chunkLoc{
			pos:  pos,
			size: int64(ch.Size),
		}

		//fmt.Printf("chunk id: <%s>\n", ch.Id)

		_, err = fr.file.Seek(c.size, os.SEEK_CUR)
		if err != nil {
			return err
		}
		pos += c.size // chunk data

		chunks[ch.Id] = &c
	}

	if pos != fileSize {
		return ErrFileFormat
	}

	// fmt
	chunkFmt, ok := chunks[token_fmt]
	if !ok {
		return ErrFileFormat
	}
	err = fr.readChunkFmt(chunkFmt)
	if err != nil {
		return err
	}

	// data
	chunkData, ok := chunks[token_data]
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
	_, err := fr.file.Seek(int64(cl.pos), os.SEEK_SET)
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
	_, err := fr.file.Seek(int64(cl.pos), os.SEEK_SET)
	if err != nil {
		return err
	}
	fr.dataLength = cl.size

	fr.r = bufio.NewReaderSize(fr.file, int(cl.size))

	return nil
}
