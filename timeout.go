package dhttp

import "time"

type Timeout time.Duration

func (_ Def) Timeout() Timeout {
	return Timeout(time.Second * 59)
}
