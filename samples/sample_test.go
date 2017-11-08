package samples

import (
	"bytes"
	"testing"
)

func TestInt16(t *testing.T) {

	bs := make([]byte, 2)

	for i := minInt16; i <= maxInt16; i++ {

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

	for i := minInt24; i <= maxInt24; i++ {

		x := int32(i)

		le.PutInt24(bs, x)
		y := le.GetInt24(bs)

		if x != y {
			t.Errorf("%d != %d", x, y)
		}
	}
}

func TestSampler8(t *testing.T) {
	var buf bytes.Buffer
	var srw sampler8
	for i := minInt8; i <= maxInt8; i++ {
		sample1 := int32(i)
		buf.Reset()
		err := srw.WriteSample(&buf, sample1)
		if err != nil {
			t.Fatal(err)
		}
		sample2, err := srw.ReadSample(&buf)
		if err != nil {
			t.Fatal(err)
		}
		if sample1 != sample2 {
			t.Fatalf("%d != %d", sample1, sample2)
		}
	}
}

func TestSampler24(t *testing.T) {
	var buf bytes.Buffer
	var srw sampler24
	for i := minInt24; i <= maxInt24; i++ {
		sample1 := int32(i)
		buf.Reset()
		err := srw.WriteSample(&buf, sample1)
		if err != nil {
			t.Fatal(err)
		}
		sample2, err := srw.ReadSample(&buf)
		if err != nil {
			t.Fatal(err)
		}
		if sample1 != sample2 {
			t.Fatalf("%d != %d", sample1, sample2)
			//t.Logf("%d != %d", sample1, sample2)
		}
	}
}
