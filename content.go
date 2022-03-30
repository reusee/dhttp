package dhttp

import (
	"io"
	"net/http"
	"time"
)

type Content []byte

func (Def) Content(
	getResp GetResponse,
	deadline RetryDeadline,
) (
	content Content,
	bs []byte,
	resp *http.Response,
) {
do:
	resp = getResp()
	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		if time.Now().Before(time.Time(deadline)) {
			goto do
		}
		ce(err)
	}
	defer resp.Body.Close()
	content = Content(bs)
	return
}

func (_ Def) StringContent(
	content []byte,
) string {
	return string(content)
}
