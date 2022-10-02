package dhttp

import "net/http"

func (Def) Transport() *http.Transport {
	// default transport
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
	}
}
