package dhttp

import "github.com/reusee/dscope"

type Headers [][2]string

var _ dscope.Reducer = Headers{}

func (_ Headers) IsReducer() {}

func (_ Def) Headers() Headers {
	return nil
}
