// +build linux

package syscall

import (
	"io"
	"os"

	"golang.org/x/sys/unix"
)

func Pwrite(fp *os.File, p []byte, offset int64) (n int, err error) {
	var m int
	for len(p) > 0 {
		m, err = unix.Pwrite(int(fp.Fd()), p, offset)
		if err != nil {
			return
		}
		if m == 0 {
			err = io.ErrUnexpectedEOF
			return
		}
		p = p[m:]
		n += m
		offset += int64(m)
	}
	return
}

func Pwritev(fp *os.File, iovs [][]byte, offset int64) (n int, err error) {
	for len(iovs) > 0 {
		i, j, m := 0, 0, 0
		m, err = unix.Pwritev(int(fp.Fd()), iovs, offset)
		if err != nil {
			return
		}
		if m == 0 {
			err = io.ErrUnexpectedEOF
			return
		}
		for i < len(iovs) && j + len(iovs[i]) <= m {
			j += len(iovs[i])
			i++
		}
		iovs = iovs[i:]
		n += j
		offset += int64(j)
	}
	return
}
