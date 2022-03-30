package dhttp

import (
	"net/http"
	"net/http/cookiejar"
)

func (_ Def) CookieJar() http.CookieJar {
	jar, err := cookiejar.New(nil)
	ce(err)
	return jar
}
