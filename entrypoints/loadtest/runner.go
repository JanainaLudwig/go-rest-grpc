package main

import "time"

type Load struct {
	CallsPerSecond uint
	Duration time.Duration
}
