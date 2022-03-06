package runner

import "time"

type Load struct {
	CallsPerSecond int
	Duration time.Duration
}
