package dhttp

import (
	"net/http"
	"time"
)

func (_ Def) Client(
	timeout Timeout,
	cookieJar http.CookieJar,
) *http.Client {
	return &http.Client{
		Timeout: time.Duration(timeout),
		Jar:     cookieJar,
	}
}
