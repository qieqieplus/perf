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
	MetricNum         = 5
	MetricSize        = 8 + 8*MetricNum
)

type Metrics struct {
	version uint32
	padding [4]byte
	data    [MetricNum]*int64
}

func cast(mem []byte) *Metrics {
	buf := mem[:MetricSize]
	ver := binary.BigEndian.Uint32(buf[:4])
	if ver != Magic {
		return nil
	}

	var data [MetricNum]*int64
	for i := 0; i < MetricNum; i++ {
		data[i] = (*int64)(unsafe.Pointer(&buf[8*(i+1)]))
	}

	return &Metrics{
		version: ver,
		padding: [4]byte{},
		data:    data,
	}
}

func initFile() (*os.File, error) {
	f, err := os.OpenFile(Filename, os.O_RDWR|os.O_CREATE, 0600|os.ModeExclusive)
	if err != nil {
		return nil, err
	}
	stat, _ := f.Stat()
	if sz := stat.Size(); sz > 0 {
		if sz != MetricSize {
			panic("init: invalid file")
		}
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

func mapMetrics(fd uintptr) []byte {
	mem, err := syscall.Mmap(int(fd), 0, MetricSize, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		panic(err)
	}
	return mem
}

func unmapMetrics(mem []byte) {
	if err := syscall.Munmap(mem); err != nil {
		panic(err)
	}
}
