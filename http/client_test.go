package http

import (
	"bytes"
	"io"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"github.com/valyala/fasthttp"
)

const (
	limitSize = 1 << 10
)

var (
	input *rand.Rand
)

func init() {
	input = rand.New(rand.NewSource(time.Now().UnixNano()))
	go func() {
		http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
			io.Copy(w, r.Body)
			return
		})
		http.ListenAndServe(":56789", nil)
	}()
}

func BenchmarkFastHttp(b *testing.B) {
	buf := &bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		func() {
			buf.Reset()
			buf.ReadFrom(io.LimitReader(input, int64(limitSize)))

			req := fasthttp.AcquireRequest()
			resp := fasthttp.AcquireResponse()
			defer func() {
				fasthttp.ReleaseRequest(req)
				fasthttp.ReleaseResponse(resp)
			}()
			req.Header.SetMethod(fasthttp.MethodPost)
			req.SetRequestURI("http://127.0.0.1:56789/echo")
			req.SetBody(buf.Bytes())
			err := fastClient.Do(req, resp)
			if err != nil {
				b.Fail()
			}
		}()
	}
}

func BenchmarkNetHttp(b *testing.B) {
	buf := &bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		func() {
			buf.Reset()
			buf.ReadFrom(io.LimitReader(input, int64(limitSize)))

			req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:56789/echo", buf)
			if err != nil {
				b.Fail()
				return
			}
			resp, err := httpClient.Do(req)
			if err != nil {
				b.Fail()
				return
			}
			resp.Body.Close()
		}()
	}
}
