package main

import (
	runner2 "grpc-rest/entrypoints/loadtest/runner"
	"time"
)

func main() {
	runner := runner2.NewRunnerGrpc("localhost:9000")
	runner.AddLoad(runner2.Load{
		CallsPerSecond: 1000,
		Duration:       15 * time.Second,
	})
	runner.AddLoad(runner2.Load{
		CallsPerSecond: 800,
		Duration:       3 * time.Second,
	})
	runner.AddLoad(runner2.Load{
		CallsPerSecond: 1000,
		Duration:       20 * time.Second,
	})
	runner.Run()
}
