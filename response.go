package dhttp

import (
	"net/http"
	"time"
)

type GetResponse func() *http.Response

func (_ Def) Response(
	client *http.Client,
	newReq NewRequest,
	deadline RetryDeadline,
) GetResponse {
	return func() *http.Response {
	do:
		resp, err := client.Do(newReq())
		if err != nil {
			if time.Now().Before(time.Time(deadline)) {
				goto do
			}
			ce(err)
		}
		return resp
	}
}
