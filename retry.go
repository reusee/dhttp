package dhttp

type Retry int

func (_ Def) Retry() Retry {
	return 1
}
