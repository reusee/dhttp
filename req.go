package dhttp

import (
	"context"

	"github.com/reusee/dscope"
)

type Method string

type Addr string

func Req(ctx context.Context, method Method, addr Addr) dscope.Scope {
	return dscope.New(&ctx, &method, &addr).Fork(Defs...)
}
