package dhttp

import (
	"context"
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
	userAgent UserAgent,
	cookie Cookie,
	headers Headers,
	ctx context.Context,
) *http.Request {

	req, err := http.NewRequestWithContext(ctx, string(method), string(addr), body)
	ce(err)

	req.Header.Add("User-Agent", string(userAgent))

	if cookie != "" {
		req.Header.Add("Cookie", string(cookie))
	}

	for _, header := range headers {
		req.Header.Add(header[0], header[1])
	}

	return req
}
