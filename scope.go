package dhttp

import "github.com/reusee/dscope"

type Def struct{}

var Defs = dscope.Methods(new(Def))
