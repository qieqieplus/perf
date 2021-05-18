package leveldb

import (
	. "github.com/qieqieplus/perf/kvdb/engine"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type Store struct {
	db *leveldb.DB
	ro *opt.ReadOptions
	wo *opt.WriteOptions
}

func NewLevelDB() Engine {
	return &Store{}
}

func (s *Store) Open(dir string) (err error) {
	s.ro = &opt.ReadOptions{}
	s.wo = &opt.WriteOptions{}

	s.db, err = leveldb.OpenFile(dir,
		&opt.Options{
			BlockCacheCapacity: 2 * 1024 * 1024,
			WriteBuffer:        2 * 1024 * 1024,
		})
	return
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
