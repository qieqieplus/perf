package leveldb

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"

	. "github.com/qieqieplus/perf/kvdb/engine"
)

type Store struct {
	db *leveldb.DB
	ro *opt.ReadOptions
	wo *opt.WriteOptions
}

func New() Engine {
	return &Store{}
}

func (s *Store) Open(dir string, options Options) (err error) {
	s.ro = &opt.ReadOptions{}
	s.wo = &opt.WriteOptions{}

	o := &opt.Options{
		BlockCacheCapacity: options.BlockCacheSize,
		WriteBuffer:        options.Memtable,
	}
	if options.BloomFilter.BitsPerKey > 0 {
		o.Filter = filter.NewBloomFilter(options.BloomFilter.BitsPerKey)
	}
	s.db, err = leveldb.OpenFile(dir, o)
	return
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) Get(key []byte) []byte {
	data, err := s.db.Get(key, s.ro)
	if err != nil {
		if err == leveldb.ErrNotFound {
			return nil
		}
		panic(err)
	}
	return data
}

func (s *Store) Put(key, value []byte) {
	err := s.db.Put(key, value, s.wo)
	if err != nil {
		panic(err)
	}
}
