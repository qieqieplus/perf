package pebble

import (
	"github.com/cockroachdb/pebble"
	"github.com/cockroachdb/pebble/bloom"
	. "github.com/qieqieplus/perf/kvdb/engine"
)

type Store struct {
	db *pebble.DB
	wo *pebble.WriteOptions
}

func New() Engine {
	return &Store{}
}

func (s *Store) Open(dir string, options Options) (err error) {
	cache := pebble.NewCache(int64(options.BlockCacheSize))
	defer cache.Unref()

	o := &pebble.Options{
		MemTableSize: options.Memtable,
		Cache:        cache,
		BytesPerSync: 2 * options.Memtable,
	}

	if options.BloomFilter.BitsPerKey > 0 {
		o.Levels = []pebble.LevelOptions{{
			FilterPolicy: bloom.FilterPolicy(8),
		}}
	}

	s.wo = &pebble.WriteOptions{Sync: false}
	s.db, err = pebble.Open(dir, o)
	return
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) Get(key []byte) []byte {
	v, closer, err := s.db.Get(key)
	defer closer.Close()

	if err != nil {
		if err == pebble.ErrNotFound {
			return nil
		}
		panic(err)
	}
	// should avoid alloc & copy here for performance
	data := make([]byte, len(v))
	copy(data, v)
	return data
}

func (s *Store) Put(key, value []byte) {
	err := s.db.Set(key, value, s.wo)
	if err != nil {
		panic(err)
	}
}
