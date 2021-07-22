package engine

var (
	Engines []Engine
)

type Options struct {
	Memtable       int
	BlockCacheSize int
	BloomFilter    struct {
		BitsPerKey int
	}
}

type Engine interface {
	Open(dir string, options Options) error
	Close()
	Get([]byte) []byte
	Put([]byte, []byte)
}
