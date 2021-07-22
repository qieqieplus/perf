package filter

import (
	"github.com/qieqieplus/perf/kvdb/engine"
	"math/rand"
	"testing"
)

func BenchmarkLevelDBGet(b *testing.B) {
	if err := ldb.Open("/tmp/leveldb", engine.Options{
		Memtable:       1 * 1024 * 1024,
		BlockCacheSize: 8 * 1024 * 1024,
		BloomFilter: struct {
			BitsPerKey int
		}{0},
	}); err != nil {
		b.Fatal(err)
		return
	}
	defer ldb.Close()

	key := make([]byte, 16)
	for i := 0; i < b.N; i++ {
		rand.Read(key)
		_ = ldb.Get(key)
	}
}

func BenchmarkLevelDBGetWithBloomFilter(b *testing.B) {
	if err := ldb.Open("/tmp/leveldb", engine.Options{
		Memtable:       1 * 1024 * 1024,
		BlockCacheSize: 8 * 1024 * 1024,
		BloomFilter: struct {
			BitsPerKey int
		}{8},
	}); err != nil {
		b.Fatal(err)
		return
	}
	defer ldb.Close()

	key := make([]byte, 16)
	for i := 0; i < b.N; i++ {
		rand.Read(key)
		_ = ldb.Get(key)
	}
}
