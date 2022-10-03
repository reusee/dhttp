package dhttp

import (
	"net/http"
	"time"
)

type GetResponse func() *http.Response

func (Def) Response(
	client *http.Client,
	newReq NewRequest,
	deadline RetryDeadline,
	maxRetry MaxRetry,
) GetResponse {
	return func() *http.Response {
	do:
		resp, err := client.Do(newReq())
		if err != nil {
			if time.Now().Before(time.Time(deadline)) {
				if maxRetry > 0 {
					maxRetry--
					goto do
				}
			}
			ce(err)
		}
		return resp
	}
}
