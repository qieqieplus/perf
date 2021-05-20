package gzip

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"math/rand"
	"testing"
	"time"
)

var (
	raw, compressed []byte
)

func init() {
	raw = make([]byte, 4*1024)
	rand.Seed(time.Now().UnixNano())
	rand.Read(raw)

	buffer := &bytes.Buffer{}
	w, _ := gzip.NewWriterLevel(buffer, gzip.BestSpeed)
	w.Write(raw)
	w.Close()
	compressed = buffer.Bytes()
}

func BenchmarkGzipNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w, _ := gzip.NewWriterLevel(ioutil.Discard, gzip.BestSpeed)
		if _, err := w.Write(raw); err != nil {
			b.Fail()
		}
		if err := w.Close(); err != nil {
			b.Fail()
		}
	}
}

func BenchmarkGzip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Gzip(ioutil.Discard, bytes.NewReader(raw))
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkGunzipNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := bytes.NewBuffer(compressed)
		rd, _ := gzip.NewReader(buf)
		io.Copy(ioutil.Discard, rd)
		rd.Close()
	}
}

func BenchmarkGunzip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := bytes.NewBuffer(compressed)
		Gunzip(ioutil.Discard, buf)
	}
}
