package main

import (
	"flag"
	runner "grpc-rest/entrypoints/loadtest/runner"
	"log"
	"time"
)

func main() {
	method := flag.String("type", "", "Specify test method: rest or grpc")
	flag.Parse()

	loads := []runner.Load{
		{
			CallsPerSecond: 50,
			Duration:       10 * time.Second,
		},
		{
			CallsPerSecond: 100,
			Duration:       10 * time.Second,
		},
		{
			CallsPerSecond: 300,
			Duration:       20 * time.Second,
		},
	}

	var test *runner.Runner
	switch *method {
	case "rest":
		test = runner.NewRunnerRest("http://localhost:8080/students", loads...)
	case "grpc":
		test = runner.NewRunnerGrpc("localhost:9000", loads...)
	default:
		log.Fatalln("Please provide a valid test method")
	}

	test.Run()

	for {
	}
}
