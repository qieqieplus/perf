package basic

import (
	"bytes"
	"os"
	"testing"

	"github.com/qieqieplus/perf/kvdb/pebble"
)

func init() {
	os.RemoveAll("/tmp/pebble")
	pbl = pebble.New()
	pbl.Open("/tmp/pebble", opts)
}

func BenchmarkPebblePut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(dict); j++ {
			pbl.Put(dict[j].key, dict[j].value)
		}
	}
}

func BenchmarkPebbleGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(dict); j++ {
			v := pbl.Get(dict[j].key)
			if !bytes.Equal(v, dict[j].value) {
				b.Fail()
			}
		}
	}
}
