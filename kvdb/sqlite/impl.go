package sqlite

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"

	. "github.com/qieqieplus/perf/kvdb/engine"
)

type Store struct {
	db *sql.DB
}

const (
	tableName = "default"
	disableSync = "PRAGMA synchronous=OFF; PRAGMA journal_mode=OFF;"
	creatTable = "CREATE TABLE IF NOT EXISTS `%s` (key VARCHAR(32), value BLOB, PRIMARY KEY (key)) WITHOUT ROWID"
	getRow = "SELECT value FROM `%s` WHERE key = ?"
	setRow = "INSERT INTO `%s` (key,value) VALUES(?,?)"
)

func New() Engine {
	return &Store{}
}

func (s *Store) Open(dir string, options Options) (err error) {
	s.db, err = sql.Open("sqlite", dir)
	if err != nil {
		s.db.Close()
		return
	}

	if err = s.db.Ping(); err != nil {
		return
	}
	s.disableSync()
	_, err = s.db.Exec(fmt.Sprintf(creatTable, tableName))
	return
}

func (s *Store) disableSync() {
	s.db.Exec(disableSync)
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) Get(key []byte) (data []byte) {
	res := s.db.QueryRow(fmt.Sprintf(getRow, tableName), key)
	if err := res.Scan(&data); err != nil {
		return nil
	}
	return
}

func (s *Store) Put(key, value []byte) {
	_, err := s.db.Exec(fmt.Sprintf(setRow, tableName), key, value)
	if err != nil {
		return
	}
}