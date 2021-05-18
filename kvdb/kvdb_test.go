package kvdb

import (
	"bytes"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/qieqieplus/perf/kvdb/engine"
	"github.com/qieqieplus/perf/kvdb/leveldb"
	"github.com/qieqieplus/perf/kvdb/pebble"
)

type pair struct {
	key   []byte
	value []byte
}

var (
	dict = make([]*pair, 1<<14)
	ldb, pdb, rdb engine.Engine
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

	os.RemoveAll("/tmp/leveldb")
	os.RemoveAll("/tmp/pebble")

	ldb  = leveldb.NewLevelDB()
	pdb  = pebble.NewPebble()

	ldb.Open("/tmp/leveldb")
	pdb.Open("/tmp/pebble")
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

func BenchmarkPebblePut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(dict); j++ {
			pdb.Put(dict[j].key, dict[j].value)
		}
	}
}

func BenchmarkPebbleGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(dict); j++ {
			v := pdb.Get(dict[j].key)
			if !bytes.Equal(v, dict[j].value) {
				b.Fail()
			}
		}
	}
}
