package route

import (
	"net/http"

	"github.com/valyala/fasthttp"
)

type point struct {
	path   string
	method string
}

type RequestHandler fasthttp.RequestHandler

var route map[point]RequestHandler

func init() {
	route = make(map[point]RequestHandler)
}

func add(path string, method string, fn RequestHandler) {
	route[point{path, method}] = fn
}

func Get(path string, fn RequestHandler) {
	add(path, http.MethodGet, fn)
}

func Post(path string, fn RequestHandler) {
	add(path, http.MethodPost, fn)
}

func Route(path string, method string) RequestHandler {
	return route[point{path, method}]
}
