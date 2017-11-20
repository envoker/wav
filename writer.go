package wav

import (
	"encoding/binary"
	"os"
)

type FileWriter struct {
	config     Config
	dataLength uint32
	file       *os.File
}

func NewFileWriter(filename string, config Config) (*FileWriter, error) {

	if err := config.checkError(); err != nil {
		return nil, err
	}

	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	fw := &FileWriter{
		config:     config,
		dataLength: 0,
		file:       file,
	}

	fw.writeConfig()

	return fw, nil
}

func (fw *FileWriter) Close() error {

	if fw.file == nil {
		return ErrFileWriterClosed
	}

	fw.writeDataLength()

	err := fw.file.Close()
	fw.file = nil
	fw.dataLength = 0
	return err
}

func (fw *FileWriter) Write(data []byte) (n int, err error) {
	n, err = fw.file.Write(data)
	fw.dataLength += uint32(n)
	return n, err
}

func (fw *FileWriter) writeConfig() error {

	_, err := fw.file.Seek(0, os.SEEK_SET)
	if err != nil {
		return err
	}

	//----------------------------------------
	// RIFF header
	ch := chunkHeader{
		Id:   tag_RIFF,
		Size: 0,
	}
	err = binary.Write(fw.file, binary.LittleEndian, ch)
	if err != nil {
		return err
	}

	//----------------------------------------
	// format WAVE
	var format = tag_WAVE
	err = binary.Write(fw.file, binary.LittleEndian, format)
	if err != nil {
		return err
	}

	//----------------------------------------
	// chunk fmt_
	fmt_data := configToFmtData(fw.config)
	ch = chunkHeader{
		Id:   tag_fmt_,
		Size: uint32(sizeFmtData),
	}
	err = binary.Write(fw.file, binary.LittleEndian, ch)
	if err != nil {
		return err
	}
	err = binary.Write(fw.file, binary.LittleEndian, fmt_data)
	if err != nil {
		return err
	}

	//----------------------------------------
	// chunk data
	ch = chunkHeader{
		Id:   tag_data,
		Size: 0,
	}
	err = binary.Write(fw.file, binary.LittleEndian, ch)
	if err != nil {
		return err
	}

	return nil
}

func (fw *FileWriter) writeDataLength() error {

	_, err := fw.file.Seek(0, os.SEEK_SET)
	if err != nil {
		return err
	}

	var riff_size = sizeWaveFormat /* riff format */ +
		(sizeChunkHeader + sizeFmtData) /* chunk fmt_ */ +
		(sizeChunkHeader + int(fw.dataLength)) /* chunk data */

	// RIFF header
	ch := chunkHeader{
		Id:   tag_RIFF,
		Size: uint32(riff_size),
	}
	err = binary.Write(fw.file, binary.LittleEndian, ch)
	if err != nil {
		return err
	}

	// data chunk
	pos := sizeWaveFormat /* riff format */ +
		(sizeChunkHeader + sizeFmtData) /* chunk fmt_ */
	_, err = fw.file.Seek(int64(pos), os.SEEK_CUR)
	if err != nil {
		return err
	}

	ch = chunkHeader{
		Id:   tag_data,
		Size: fw.dataLength,
	}
	err = binary.Write(fw.file, binary.LittleEndian, ch)
	if err != nil {
		return err
	}

	return nil
}
