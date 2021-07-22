package filter

import (
	"math/rand"
	"os"
	"time"

	"github.com/qieqieplus/perf/kvdb/engine"
	"github.com/qieqieplus/perf/kvdb/leveldb"
)

var (
	ldb, pbl engine.Engine
)

func init() {
	rand.Seed(time.Now().UnixNano())
	os.RemoveAll("/tmp/leveldb")

	ldb = leveldb.New()
	fillDB(ldb)
}

func fillDB(db engine.Engine) error {
	if err := db.Open("/tmp/leveldb", engine.Options{
		Memtable:       8 * 1024 * 1024,
		BlockCacheSize: 8 * 1024 * 1024,
		BloomFilter: struct {
			BitsPerKey int
		}{8},
	}); err != nil {
		return err
	}
	defer db.Close()

	key := make([]byte, 16)
	value := make([]byte, 64)

	for i := 0; i < 1e6; i++ {
		rand.Read(key)
		rand.Read(value)
		db.Put(key, value)
	}
	return nil
}
