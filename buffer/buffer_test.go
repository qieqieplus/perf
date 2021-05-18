package buffer

import (
	"bytes"
	"testing"
)

var (
	size = 1 << 20
	data = make([]byte, size)
)

func BenchmarkRead1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Read1(bytes.NewReader(data), size)
	}
}

func BenchmarkRead2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Read2(bytes.NewReader(data), size)
	}
}

func BenchmarkRead3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Read3(bytes.NewReader(data), size)
	}
}

func BenchmarkRead4(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Read4(bytes.NewReader(data), size)
	}
}
