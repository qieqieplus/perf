// +build linux

package syscall

import (
	"math/rand"
	"os"
	"testing"
	"time"
)

const (
	batch   = 1 << 8
	iter    = batch * 64
	bufSize = 1 << 12
)

var (
	buffer [][]byte
	f1, f2 *os.File
)

func init() {
	var err error
	if f1, err = os.Create("test1"); err != nil {
		panic(err)
	}
	if f2, err = os.Create("test2"); err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UnixNano())

	buffer = make([][]byte, iter)
	for i := range buffer {
		buffer[i] = make([]byte, bufSize)
		rand.Read(buffer[i])
	}
}

func BenchmarkPwrite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		offset := int64(0)
		for _, buf := range buffer {
			n, _ := Pwrite(f1, buf, offset)
			if n != len(buf) {
				b.Fail()
				return
			}
			offset += int64(n)
		}
		//f1.Truncate(0)
	}
}

func BenchmarkPwritev(b *testing.B) {
	for i := 0; i < b.N; i++ {
		offset := int64(0)
		for j := 0; j < len(buffer); j += batch {
			n, _ := Pwritev(f2, buffer[j:j+batch], offset)
			if n != bufSize*batch {
				b.Fail()
				return
			}
			offset += int64(n)
		}
		//f2.Truncate(0)
	}
}
