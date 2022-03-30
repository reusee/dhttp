package dhttp

import "net/http"

func (_ Def) Response(
	client *http.Client,
	req *http.Request,
) *http.Response {
	resp, err := client.Do(req)
	ce(err)
	return resp
}
