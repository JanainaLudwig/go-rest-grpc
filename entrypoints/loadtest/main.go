package main

import (
	runner2 "grpc-rest/entrypoints/loadtest/runner"
	"time"
)

func main() {
	runner := runner2.NewRunnerGrpc("localhost:9000")
	runner.AddLoad(runner2.Load{
		CallsPerSecond: 5,
		Duration:       1 * time.Second,
	})
	runner.AddLoad(runner2.Load{
		CallsPerSecond: 20,
		Duration:       2 * time.Second,
	})
	runner.AddLoad(runner2.Load{
		CallsPerSecond: 30,
		Duration:       3 * time.Second,
	})
	runner.Run()
}
