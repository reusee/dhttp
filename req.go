package dhttp

import (
	"context"

	"github.com/reusee/dscope"
)

type Method string

type Addr string

func Req(ctx context.Context, method string, addr string) dscope.Scope {
	return dscope.New(
		&ctx,
		func() (Method, Addr) {
			return Method(method), Addr(addr)
		},
	).Fork(Defs...)
}
