package samples

type littleEndian struct{}

var le littleEndian

func (littleEndian) PutInt8(b []byte, i int8) {
	u := uint8(int(i) + 128)
	b[0] = byte(u)
}

func (littleEndian) GetInt8(b []byte) int8 {
	u := uint8(b[0])
	return int8(int(u) - 128)
}

func (littleEndian) PutInt16(b []byte, i int16) {

	u := uint16(i)

	b[0] = byte(u >> 0)
	b[1] = byte(u >> 8)
}

func (littleEndian) GetInt16(b []byte) int16 {

	u := uint16(b[0]) << 0
	u |= uint16(b[1]) << 8

	return int16(u)
}

func (littleEndian) PutInt24(b []byte, i int32) {

	u := uint32(i)

	b[0] = byte(u >> 0)
	b[1] = byte(u >> 8)
	b[2] = byte(u >> 16)
}

func (littleEndian) GetInt24(b []byte) int32 {

	var u uint32

	u |= uint32(b[0]) << 0
	u |= uint32(b[1]) << 8
	u |= uint32(b[2]) << 16

	if (b[2] & 0x80) != 0 { // is negative
		u |= uint32(0xFF) << 24
	}

	return int32(u)
}

func (littleEndian) PutInt32(b []byte, i int32) {

	u := uint32(i)

	b[0] = byte(u >> 0)
	b[1] = byte(u >> 8)
	b[2] = byte(u >> 16)
	b[3] = byte(u >> 24)
}

func (littleEndian) GetInt32(b []byte) int32 {

	var u uint32

	u |= uint32(b[0]) << 0
	u |= uint32(b[1]) << 8
	u |= uint32(b[2]) << 16
	u |= uint32(b[3]) << 24

	return int32(u)
}
