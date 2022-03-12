package main

import (
	"github.com/newrelic/go-agent/v3/newrelic"
	"grpc-rest/api/router"
	"grpc-rest/config"
	"grpc-rest/core"
	"log"
)

func main() {

	config.LoadEnv(config.RootPath() + "/config/.env")

	log.Println("Starting the " + config.App.AppEnv + " API...", "Go to http://localhost:" + config.App.ApiPort)

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("comparison-go-rest"),
		newrelic.ConfigLicense(config.App.NewRelicLicence),
		newrelic.ConfigDistributedTracerEnabled(true),
	)
	if err != nil {
		log.Fatalln(err)
	}
	config.App.SetNewrelicApp(app)

	core.StartApp(true)

	router.StartApi()
}
