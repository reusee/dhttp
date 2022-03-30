package dhttp

import "net/http"

func (_ Def) Client() *http.Client {
	return http.DefaultClient
}
