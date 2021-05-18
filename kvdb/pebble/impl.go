package pebble

import (
	"github.com/cockroachdb/pebble"
	. "github.com/qieqieplus/perf/kvdb/engine"
)

type Store struct {
	db *pebble.DB
	wo *pebble.WriteOptions
}

func NewPebble() Engine {
	return &Store{}
}

func (s *Store) Open(dir string) (err error) {
	s.wo = &pebble.WriteOptions{}
	s.db, err = pebble.Open(dir, &pebble.Options{
		MemTableSize: 512 * 1024,
	})
	return
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