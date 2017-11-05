package sample

type littleEndian struct{}

var le littleEndian

//	Int16

func (littleEndian) PutInt16(b []byte, i int16) {

	u := uint16(i)

	b[0] = byte(u)
	b[1] = byte(u >> 8)
}

func (littleEndian) GetInt16(b []byte) (i int16) {

	u := uint16(b[0])
	u |= uint16(b[1]) << 8

	i = int16(u)

	return
}

//	Int24

func (littleEndian) PutInt24(b []byte, i int32) {

	u := uint32(i)

	b[0] = byte(u)
	b[1] = byte(u >> 8)
	b[2] = byte(u >> 16)
}

func (littleEndian) GetInt24(b []byte) (i int32) {

	u := uint32(b[0])
	u |= uint32(b[1]) << 8
	u |= uint32(b[2]) << 16

	if (b[2] & 0x80) == 0x80 {
		u |= uint32(0xFF) << 24
	}

	i = int32(u)

	return
}

//	Int32

func (littleEndian) PutInt32(b []byte, i int32) {

	u := uint32(i)

	b[0] = byte(u)
	b[1] = byte(u >> 8)
	b[2] = byte(u >> 16)
	b[3] = byte(u >> 24)
}

func (littleEndian) GetInt32(b []byte) (i int32) {

	u := uint32(b[0])
	u |= uint32(b[1]) << 8
	u |= uint32(b[2]) << 16
	u |= uint32(b[3]) << 24

	i = int32(u)

	return
}
