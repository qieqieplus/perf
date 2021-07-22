package basic

import (
	"bytes"
	"os"
	"testing"

	"github.com/qieqieplus/perf/kvdb/leveldb"
)

func init() {
	os.RemoveAll("/tmp/leveldb")
	ldb = leveldb.New()
	ldb.Open("/tmp/leveldb", opts)
}

func BenchmarkLevelDBPut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(dict); j++ {
			ldb.Put(dict[j].key, dict[j].value)
		}
	}
}

func BenchmarkLevelDBGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(dict); j++ {
			v := ldb.Get(dict[j].key)
			if !bytes.Equal(v, dict[j].value) {
				b.Fail()
			}
		}
	}
}
