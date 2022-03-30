package dhttp

import "net/http"

func (_ Def) Response(
	client *http.Client,
	req *http.Request,
	retry Retry,
) *http.Response {

do:
	resp, err := client.Do(req)
	if err != nil {
		if retry > 0 {
			retry--
			goto do
		}
		ce(err)
	}

	return resp
}
