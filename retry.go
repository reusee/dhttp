package dhttp

import "time"

type RetryDuration time.Duration

type RetryDeadline time.Time

func (Def) RetryDuration() RetryDuration {
	return 0
}

func (Def) RetryDeadline(
	d RetryDuration,
) RetryDeadline {
	return RetryDeadline(time.Now().Add(time.Duration(d)))
}

type MaxRetry int

func (Def) MaxRetry() MaxRetry {
	return 8
}
