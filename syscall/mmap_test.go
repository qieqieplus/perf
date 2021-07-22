package syscall

import (
	"testing"
)

func BenchmarkMmapWrite(b *testing.B) {
	fd, err := initFile()
	if err != nil {
		b.Fail()
		return
	}
	mem := mapMetrics(fd.Fd())
	defer unmapMetrics(mem)

	//fmt.Println(mem)
	metric := cast(mem)
	if metric == nil {
		b.Fail()
		return
	}

	for i := 0; i < b.N; i++ {
		*metric.data[0] = 0x1234567887654321
		*metric.data[4] = 0x7fffffffffffffff
	}
}
