package buffer

import "C"
import (
	"reflect"
	"unsafe"
)

var (
	cPtr  *C.uchar
	cSize C.size_t
)

func init() {
	cSize = C.size_t(1 << 16)
	cPtr = (*C.uchar)(C.malloc(cSize))
}

func ToBytes1(p *C.uchar, size C.size_t) []byte {
	return C.GoBytes(unsafe.Pointer(p), C.int(size))
}

func ToBytes2(p *C.uchar, size C.size_t) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: uintptr(unsafe.Pointer(p)),
			Len:  int(size),
			Cap:  int(size),
		},
	))
}

// ================
// use of cgo in test is not supported

func ToBytesTest1() []byte {
	return ToBytes1(cPtr, cSize)
}

func ToBytesTest2() []byte {
	return ToBytes2(cPtr, cSize)
}
