package dhttp

import "github.com/reusee/dscope"

type Headers [][2]string

var _ dscope.Reducer = Headers{}

func (Headers) IsReducer() {}

func (Def) Headers() Headers {
	return nil
}
