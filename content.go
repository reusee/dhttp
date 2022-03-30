package dhttp

import (
	"io"
	"net/http"

	"github.com/reusee/e4"
)

type Content []byte

func (Def) Content(
	resp *http.Response,
) (
	content Content,
	bs []byte,
) {
	bs, err := io.ReadAll(resp.Body)
	ce(err, e4.Close(resp.Body))
	resp.Body.Close()
	content = Content(bs)
	return
}

func (_ Def) StringContent(
	content []byte,
) string {
	return string(content)
}
