package dhttp

import (
	"net/http"
	"net/http/cookiejar"
)

func (Def) CookieJar() http.CookieJar {
	jar, err := cookiejar.New(nil)
	ce(err)
	return jar
}
