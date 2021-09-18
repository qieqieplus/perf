// +build sqlite

package basic

import (
	"bytes"
	"os"
	"testing"

	"github.com/qieqieplus/perf/kvdb/sqlite"
)

func init() {
	os.RemoveAll("/tmp/sqlite")
	lite = sqlite.New()
	lite.Open("/tmp/sqlite", opts)
}

func BenchmarkSqlitePut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(dict); j++ {
			lite.Put(dict[j].key, dict[j].value)
		}
	}
}

func BenchmarkSqliteGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(dict); j++ {
			v := lite.Get(dict[j].key)
			if !bytes.Equal(v, dict[j].value) {
				b.Fail()
			}
		}
	}
}
