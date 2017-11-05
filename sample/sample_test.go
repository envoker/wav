package sample

import (
	"testing"
)

func TestInt16(t *testing.T) {

	bs := make([]byte, 2)

	for i := -(maxInt16 + 1); i <= maxInt16; i++ {

		x := int16(i)

		le.PutInt16(bs, x)
		y := le.GetInt16(bs)

		if x != y {
			t.Errorf("%d != %d", x, y)
		}
	}
}

func TestInt24(t *testing.T) {

	bs := make([]byte, 3)

	for i := -(maxInt24 + 1); i <= maxInt24; i++ {

		x := int32(i)

		le.PutInt24(bs, x)
		y := le.GetInt24(bs)

		if x != y {
			t.Errorf("%d != %d", x, y)
		}
	}
}
