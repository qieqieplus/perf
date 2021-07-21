package syscall

import (
	"encoding/binary"
	"os"
	"syscall"
	"unsafe"
)

const (
	Filename          = "MANIFEST"
	Magic      uint32 = 0xCAC01000
	MetricSize        = 48
)

type Metrics struct {
	version       uint32
	padding       [4]byte
	a, b, c, d, e *int64
}

func cast(buf []byte) *Metrics {
	buf = buf[:MetricSize]
	ver := binary.BigEndian.Uint32(buf[:4])
	if ver != Magic {
		return nil
	}
	return &Metrics{
		version: ver,
		padding: [4]byte{},
		a:       (*int64)(unsafe.Pointer(&buf[8*1])),
		b:       (*int64)(unsafe.Pointer(&buf[8*2])),
		c:       (*int64)(unsafe.Pointer(&buf[8*3])),
		d:       (*int64)(unsafe.Pointer(&buf[8*4])),
		e:       (*int64)(unsafe.Pointer(&buf[8*5])),
	}
}

func initFile() (*os.File, error) {
	f, err := os.OpenFile(Filename, os.O_RDWR|os.O_CREATE, 0600|os.ModeExclusive)
	if err != nil {
		return nil, err
	}
	stat, _ := f.Stat()
	if stat.Size() > 0 {
		return f, nil
	}
	buf := make([]byte, MetricSize)
	binary.BigEndian.PutUint32(buf[:4], Magic)

	_, err = f.WriteAt(buf, 0)
	if err != nil {
		return nil, err
	}
	f.Sync()
	return f, nil
}

func mapMetrics() *Metrics {
	fd, err := initFile()
	if err != nil {
		panic(err)
	}
	mem, err := syscall.Mmap(int(fd.Fd()), 0, MetricSize, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	//fmt.Println(mem)
	if err != nil {
		panic(err)
	}
	data := cast(mem)
	if data == nil {
		panic("mmap: invalid file")
	}
	return data
}
