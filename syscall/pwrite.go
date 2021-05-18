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
	var vl, vs int
	if vl = len(iovs); vl == 0 {
		return
	} else if vs = len(iovs[0]); vs == 0 {
		return
	}

	var m int
	for len(iovs) > 0 {
		m, err = unix.Pwritev(int(fp.Fd()), iovs, offset)
		if err != nil {
			return
		}
		if m == 0 {
			err = io.ErrUnexpectedEOF
			return
		}
		m -= m % vs
		iovs = iovs[m/vs:]
		n += m
		offset += int64(m)
	}
	return
}
