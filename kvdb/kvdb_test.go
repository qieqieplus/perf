package kvdb

import (
	"bytes"
	"math/rand"
	"testing"
	"time"

	"github.com/qieqieplus/perf/kvdb/leveldb"
	"github.com/qieqieplus/perf/kvdb/pebble"
)

type pair struct {
	key   []byte
	value []byte
}

var (
	dict = make([]*pair, 1<<12)
	ldb  = leveldb.NewLevelDB()
	pdb  = pebble.NewPebble()
)

func init() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < cap(dict); i++ {
		k := make([]byte, 1<<4)
		v := make([]byte, 1<<8)

		rand.Read(k)
		rand.Read(v)

		dict[i] = &pair{k, v}
	}

	var err error
	err = ldb.Open("/tmp/leveldb")
	if err != nil {
		panic(err)
	}
	err = pdb.Open("/tmp/pebble")
	if err != nil {
		panic(err)
	}
}

func BenchmarkLevelDBPut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < cap(dict); j++ {
			ldb.Put(dict[j].key, dict[j].value)
		}
	}
}

func BenchmarkLevelDBGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < cap(dict); j++ {
			v := ldb.Get(dict[j].key)
			if !bytes.Equal(v, dict[j].value) {
				b.Fail()
			}
		}
	}
}

func BenchmarkPebblePut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < cap(dict); j++ {
			pdb.Put(dict[j].key, dict[j].value)
		}
	}
}

func BenchmarkPebbleGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < cap(dict); j++ {
			v := pdb.Get(dict[j].key)
			if !bytes.Equal(v, dict[j].value) {
				b.Fail()
			}
		}
	}
}
