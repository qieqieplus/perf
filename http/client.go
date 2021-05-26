package http

import (
	"net/http"

	"github.com/valyala/fasthttp"
)

var (
	httpClient *http.Client
	fastClient *fasthttp.Client
)

func init() {
	httpClient = &http.Client{
		Transport: &http.Transport{
			MaxConnsPerHost: 1024,
		}}

	fastClient = &fasthttp.Client{
		MaxConnsPerHost: 1024,
	}
}
