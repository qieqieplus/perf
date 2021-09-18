package basic

import (
	"math/rand"
	"time"

	"github.com/qieqieplus/perf/kvdb/engine"
)

type pair struct {
	key   []byte
	value []byte
}

var (
	dict = make([]*pair, 1<<12)
	opts = engine.Options{
		Memtable:       1 * 1024 * 1024,
		BlockCacheSize: 8 * 1024 * 1024,
		BloomFilter: struct {
			BitsPerKey int
		}{8},
	}
	ldb, pbl, rdb, lite engine.Engine
)

func init() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(dict); i++ {
		k := make([]byte, 1<<4)
		v := make([]byte, 1<<7)

		rand.Read(k)
		rand.Read(v)

		dict[i] = &pair{k, v}
	}
}
