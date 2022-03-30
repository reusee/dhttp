package dhttp

import (
	"io"
	"net/http"
)

type RequestBody interface {
	io.Reader
}

func (_ Def) RequestBody() RequestBody {
	return nil
}

func (_ Def) Request(
	method Method,
	addr Addr,
	body RequestBody,
) *http.Request {
	req, err := http.NewRequest(string(method), string(addr), body)
	ce(err)
	return req
}
