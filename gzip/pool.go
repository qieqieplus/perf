package gzip

import (
	"io"
	"sync"
)

var (
	copyBufferPool = NewBytesPool(32 * 1024)
)

func NewBytesPool(size int) sync.Pool {
	return sync.Pool{
		New: func() interface{} {
			return make([]byte, 0, size)
		},
	}
}

func Copy(dst io.Writer, src io.Reader) (written int64, err error) {
	buf := copyBufferPool.Get().([]byte)
	defer copyBufferPool.Put(buf)
	return io.CopyBuffer(dst, src, buf[:cap(buf)])
}
