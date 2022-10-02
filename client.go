package dhttp

import (
	"net/http"
	"time"
)

func (Def) Client(
	timeout Timeout,
	cookieJar http.CookieJar,
	transport *http.Transport,
) *http.Client {
	return &http.Client{
		Transport: transport,
		Timeout:   time.Duration(timeout),
		Jar:       cookieJar,
	}
}
