package main

import (
	runner "grpc-rest/entrypoints/loadtest/runner"
	"time"
)

func main() {
	loads := []runner.Load{
		{
			CallsPerSecond: 700,
			Duration:       3 * time.Second,
		},
		{
			CallsPerSecond: 800,
			Duration:       3 * time.Second,
		},
		{
			CallsPerSecond: 1000,
			Duration:       3 * time.Second,
		},
	}

	//runnerGrpc := runner.NewRunnerGrpc("localhost:9000", loads...)
	//runnerGrpc.Run()

	runnerGrpc := runner.NewRunnerRest("http://localhost:8080/students", loads...)
	runnerGrpc.Run()
}
