// +build rocksdb

package kvdb

import (
	"bytes"
	"os"
	"testing"

	"github.com/qieqieplus/perf/kvdb/rocksdb"
)

func init() {
	os.RemoveAll("/tmp/rocksdb")
	rdb = rocksdb.NewRocksDB()
	rdb.Open("/tmp/rocksdb")
}

func BenchmarkRocksDBPut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(dict); j++ {
			rdb.Put(dict[j].key, dict[j].value)
		}
	}
}

func BenchmarkRocksDBGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(dict); j++ {
			v := rdb.Get(dict[j].key)
			if !bytes.Equal(v, dict[j].value) {
				b.Fail()
			}
		}
	}
}
