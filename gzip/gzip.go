package gzip

import (
	"compress/gzip"
	"fmt"
	"io"
	"sync"
)

var (
	zipReaderPool = sync.Pool{
		New: func() interface{} {
			return new(gzip.Reader)
			//return gzip.NewReader()
		},
	}

	zipWriterPool = sync.Pool{
		New: func() interface{} {
			w, _ := gzip.NewWriterLevel(nil, gzip.BestSpeed)
			return w
		},
	}
)

func Gzip(w io.Writer, r io.Reader) (int64, error) {
	gw, ok := zipWriterPool.Get().(*gzip.Writer)
	if !ok {
		return 0, fmt.Errorf("gzip: new writer error")
	}
	gw.Reset(w)
	defer func() {
		gw.Close()
		zipWriterPool.Put(gw)
	}()
	return Copy(gw, r)
}

func Gunzip(w io.Writer, r io.Reader) (int64, error) {
	gr, ok := zipReaderPool.Get().(*gzip.Reader)
	if !ok {
		return 0, fmt.Errorf("gzip: new reader error")
	}
	err := gr.Reset(r)
	if err != nil {
		return 0, err
	}
	defer func() {
		gr.Close()
		zipReaderPool.Put(gr)
	}()
	return Copy(w, gr)
}
