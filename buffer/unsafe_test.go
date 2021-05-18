package buffer

import (
	"bytes"
	"testing"
)

func TestToBytes(t *testing.T) {
	b1 := ToBytesTest1()
	b2 := ToBytesTest2()

	if !bytes.Equal(b1, b2) {
		t.Fail()
	}
}

func BenchmarkToBytes1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := ToBytesTest1()
		if len(s) != int(cSize) {
			b.Fail()
		}
	}
}

func BenchmarkToBytes2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := ToBytesTest2()
		if len(s) != int(cSize) {
			b.Fail()
		}
	}
}
