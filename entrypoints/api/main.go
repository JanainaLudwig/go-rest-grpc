package main

import (
	"grpc-rest/api/router"
	"grpc-rest/config"
	"grpc-rest/core"
	"log"
)

func main() {

	config.LoadEnv(config.RootPath() + "/config/.env")

	log.Println("Starting the " + config.App.AppEnv + " API...", "Go to http://localhost:" + config.App.ApiPort)

	core.StartApp()

	router.StartApi()
}
