package dhttp

import (
	"context"
	"io"
	"net/http"
)

type NewRequestBody func() io.Reader

func (_ Def) RequestBody() NewRequestBody {
	return func() io.Reader {
		return nil
	}
}

type NewRequest func() *http.Request

func (_ Def) Request(
	method Method,
	addr Addr,
	newReqBody NewRequestBody,
	userAgent UserAgent,
	cookie Cookie,
	headers Headers,
	ctx context.Context,
) NewRequest {

	return func() *http.Request {
		req, err := http.NewRequestWithContext(ctx, string(method), string(addr), newReqBody())
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

}
