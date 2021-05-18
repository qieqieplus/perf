package protobuf

import (
	"math/rand"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/qieqieplus/perf/protobuf/gogo"
	"github.com/qieqieplus/perf/protobuf/golang"
)

var (
	link    = make([]byte, 1<<8)
	counter = rand.Int31()

	goMarshal   []byte
	gogoMarshal []byte
)

func init() {
	rand.Seed(time.Now().UnixNano())

	_, err := rand.Read(link)
	if err != nil {
		panic(err)
	}
}

func BenchmarkMarshalGolang(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		entry := &golang.Entry{
			Name:        "Benchmark",
			Extended:    make(map[string][]byte),
			Link:        link,
			LinkCounter: counter,
		}
		goMarshal, err = proto.Marshal(entry)
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkMarshalGogo(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		entry := &gogo.Entry{
			Name:        "Benchmark",
			Extended:    make(map[string][]byte),
			Link:        link,
			LinkCounter: counter,
		}
		gogoMarshal, err = entry.Marshal()
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkUnmarshalGolang(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		entry := &golang.Entry{}
		err = proto.Unmarshal(goMarshal, entry)
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkUnmarshalGogo(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		entry := &gogo.Entry{}
		err = entry.Unmarshal(gogoMarshal)
		if err != nil {
			b.Fail()
		}
	}
}
