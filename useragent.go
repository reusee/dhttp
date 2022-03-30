package dhttp

type UserAgent string

func (_ Def) UserAgent() UserAgent {
	return "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:98.0) Gecko/20100101 Firefox/98.0"
}
