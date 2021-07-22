package rocksdb

import (
	"github.com/qieqieplus/gorocksdb/v5"
	. "github.com/qieqieplus/perf/kvdb/engine"
)

type Store struct {
	db    *gorocksdb.DB
	cache *gorocksdb.Cache

	o   *gorocksdb.Options
	bto *gorocksdb.BlockBasedTableOptions
	wo  *gorocksdb.WriteOptions
	ro  *gorocksdb.ReadOptions
}

func New() Engine {
	return &Store{
		o:   gorocksdb.NewDefaultOptions(),
		bto: gorocksdb.NewDefaultBlockBasedTableOptions(),
		wo:  gorocksdb.NewDefaultWriteOptions(),
		ro:  gorocksdb.NewDefaultReadOptions(),
	}
}

func (s *Store) Open(dir string, options Options) (err error) {
	s.cache = gorocksdb.NewLRUCache(uint64(options.BlockCacheSize))

	s.bto.SetBlockCache(s.cache)
	if options.BloomFilter.BitsPerKey > 0 {
		s.bto.SetFilterPolicy(gorocksdb.NewBloomFilterFull(options.BloomFilter.BitsPerKey))
	}

	s.o.SetWriteBufferSize(options.Memtable)
	s.o.SetBlockBasedTableFactory(s.bto)
	s.o.SetCreateIfMissing(true)

	s.db, err = gorocksdb.OpenDb(s.o, dir)
	return
}

func (s *Store) Close() {
	s.o.Destroy()
	s.bto.Destroy()
	s.wo.Destroy()
	s.ro.Destroy()

	s.cache.Destroy()
	s.db.Close()
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
