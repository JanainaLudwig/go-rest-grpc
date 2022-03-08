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
			CallsPerSecond: 500,
			Duration:       3 * time.Second,
		},
		{
			CallsPerSecond: 700,
			Duration:       3 * time.Second,
		},
		{
			CallsPerSecond: 800,
			Duration:       3 * time.Second,
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
