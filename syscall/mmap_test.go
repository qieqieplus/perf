package syscall

import (
	"testing"
)

func BenchmarkMmapWrite(b *testing.B) {
	m := mapMetrics()
	for i := 0; i < b.N; i++ {
		*m.a = 0x1234567887654321
		*m.d = 0x7fffffffffffffff
	}
}
