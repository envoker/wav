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

func NewFileWriter(fileName string, config Config) (*FileWriter, error) {

	if err := config.checkError(); err != nil {
		return nil, err
	}

	file, err := os.Create(fileName)
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
	riff_h := riffHeader{
		Id:     token_RIFF,
		Size:   0,
		Format: token_WAVE,
	}
	err = binary.Write(fw.file, binary.LittleEndian, riff_h)
	if err != nil {
		return err
	}

	//----------------------------------------
	// fmt chunk
	fmt_h := chunkHeader{
		Id:   token_fmt,
		Size: sizeofFmtData,
	}
	err = binary.Write(fw.file, binary.LittleEndian, fmt_h)
	if err != nil {
		return err
	}
	fmt_data := configToFmtData(fw.config)
	err = binary.Write(fw.file, binary.LittleEndian, fmt_data)
	if err != nil {
		return err
	}

	//----------------------------------------
	// data chunk
	data_h := chunkHeader{
		Id:   token_data,
		Size: 0,
	}
	err = binary.Write(fw.file, binary.LittleEndian, data_h)
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

	size_fmtChunk := uint32(sizeofChunkHeader + sizeofFmtData)
	size_dataChunk := uint32(sizeofChunkHeader + fw.dataLength)
	riff_size := sizeofRiffFormat + size_fmtChunk + size_dataChunk

	// RIFF header
	riff_h := riffHeader{
		Id:     token_RIFF,
		Size:   riff_size,
		Format: token_WAVE,
	}
	err = binary.Write(fw.file, binary.LittleEndian, riff_h)
	if err != nil {
		return err
	}

	// data chunk
	_, err = fw.file.Seek(int64(size_fmtChunk), os.SEEK_CUR)
	if err != nil {
		return err
	}

	data_h := chunkHeader{
		Id:   token_data,
		Size: fw.dataLength,
	}
	err = binary.Write(fw.file, binary.LittleEndian, data_h)
	if err != nil {
		return err
	}

	return nil
}
