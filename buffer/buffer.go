package buffer

import (
	"bytes"
	"io"
	"io/ioutil"
)

func Read1(r io.Reader, size int) []byte {
	ret, _ := ioutil.ReadAll(r)
	return ret
}

func Read2(r io.Reader, size int) []byte {
	var buf bytes.Buffer
	buf.ReadFrom(r)
	return buf.Bytes()
}

func Read3(r io.Reader, size int) []byte {
	var buf bytes.Buffer
	buf.Grow(size)
	buf.ReadFrom(r)
	return buf.Bytes()
}

func Read4(r io.Reader, size int) []byte {
	var buf bytes.Buffer
	buf.Grow(size + bytes.MinRead)
	buf.ReadFrom(r)
	return buf.Bytes()
}
