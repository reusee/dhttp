package dhttp

import "time"

type RetryDuration time.Duration

type RetryDeadline time.Time

func (_ Def) RetryDuration() RetryDuration {
	return 0
}

func (_ Def) RetryDeadline(
	d RetryDuration,
) RetryDeadline {
	return RetryDeadline(time.Now().Add(time.Duration(d)))
}
