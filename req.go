package dhttp

import "github.com/reusee/dscope"

type Method string

type Addr string

func Req(method Method, addr Addr) dscope.Scope {
	return dscope.New(&method, &addr).Fork(Defs...)
}
