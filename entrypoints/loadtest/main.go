package main

import (
	"encoding/json"
	"flag"
	"grpc-rest/config"
	"grpc-rest/entrypoints/loadtest/runner"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	config.LoadEnv(config.RootPath() + "/config/.env")

	method := flag.String("type", "", "Specify test method: rest or grpc")
	methodType := flag.String("method", "get", "Specify service method: get or post")
	flag.Parse()

	loads := getLoadConfig()

	//loads := []runner.Load{
	//	{CallsPerSecond: 2, Duration: 1 * time.Second},
	//	{CallsPerSecond: 3, Duration: 1 * time.Second},
	//}

	testRunner := getTestRunner(method, loads, *methodType)

	report := testRunner.Run(60)
	testRunner.ReportToCsv()

	log.Println(report)
}

type loadConfig struct {
	Calls   int `json:"calls"`
	Seconds int `json:"seconds"`
}

func getLoadConfig() []runner.Load {
	plan, err := ioutil.ReadFile(config.RootPath() + "/config/load.json")
	if err != nil {
		log.Fatalln(err)
	}
	var data []loadConfig
	err = json.Unmarshal(plan, &data)
	if err != nil {
		log.Fatalln(err)
	}

	var loadRun []runner.Load
	for _, load := range data {
		loadRun = append(loadRun, runner.Load{
			CallsPerSecond: load.Calls,
			Duration:       time.Duration(load.Seconds) * time.Second,
		})
	}

	return loadRun
}

func getTestRunner(method *string, loads []runner.Load, methodType string) *runner.Runner {
	var test *runner.Runner
	switch *method {
	case "rest":
		test = runner.NewRunnerRest(config.App.ServerRest+"/students", methodType, loads...)
	case "grpc":
		test = runner.NewRunnerGrpc(config.App.ServerGrpc, methodType, loads...)
	default:
		log.Fatalln("Please provide a valid test method")
	}

	return test
}
