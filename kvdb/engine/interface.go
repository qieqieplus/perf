package engine

type Engine interface {
	Open(dir string) error
	Get([]byte) []byte
	Put([]byte, []byte)
}
