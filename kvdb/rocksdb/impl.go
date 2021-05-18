
package rocksdb

import (
	"github.com/qieqieplus/gorocksdb/v5"
	. "github.com/qieqieplus/perf/kvdb/engine"
)


type Store struct {
	db *gorocksdb.DB
	o *gorocksdb.Options
	wo *gorocksdb.WriteOptions
	ro *gorocksdb.ReadOptions
}

func NewRocksDB() Engine {
	return &Store{
		o: gorocksdb.NewDefaultOptions(),
		wo: gorocksdb.NewDefaultWriteOptions(),
		ro: gorocksdb.NewDefaultReadOptions(),
	}
}

func (s *Store) Open(dir string) (err error) {
	s.o.SetCreateIfMissing(true)
	s.o.SetWriteBufferSize(2 * 1024 * 1024)
	s.db, err = gorocksdb.OpenDb(s.o, dir)
	return
}

func (s *Store) Get(key []byte) []byte {
	slice, err := s.db.Get(s.ro, key)
	if err != nil {
		if slice == nil || !slice.Exists() {
			return nil
		}
		panic(err)
	}
	defer slice.Free()
	// should avoid alloc & copy here for performance
	data := make([]byte, slice.Size())
	copy(data, slice.Data())
	return data
}

func (s *Store) Put(key, value []byte) {
	err := s.db.Put(s.wo, key, value)
	if err != nil {
		panic(err)
	}
}
