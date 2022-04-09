package main

import (
	"flag"
	"grpc-rest/entrypoints/loadtest/runner"
	"log"
	"time"
)

func main() {
	method := flag.String("type", "", "Specify test method: rest or grpc")
	flag.Parse()

	loads := []runner.Load{
		{CallsPerSecond: 10, Duration: 2 * time.Second},
		{CallsPerSecond: 20, Duration: 2 * time.Second},
	}

	testRunner := getTestRunner(method, loads)

	report := testRunner.Run()
	testRunner.ReportToCsv()

	log.Println(report)
}

func getTestRunner(method *string, loads []runner.Load) *runner.Runner {
	var test *runner.Runner
	switch *method {
	case "rest":
		test = runner.NewRunnerRest("http://localhost:8080/students", loads...)
	case "grpc":
		test = runner.NewRunnerGrpc("localhost:9000", loads...)
	default:
		log.Fatalln("Please provide a valid test method")
	}

	return test
}
