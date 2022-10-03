package dhttp

import "time"

type Timeout time.Duration

func (Def) Timeout() Timeout {
	return Timeout(time.Second * 59)
}
