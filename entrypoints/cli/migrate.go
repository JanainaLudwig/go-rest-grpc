package main

import (
	"flag"
	"fmt"
	"grpc-rest/config"
	"grpc-rest/core"
	"log"
	"os"
	"time"
)

func main() {
	config.LoadEnv(config.RootPath() + "/config/.env")
	core.StartApp()

	action := flag.String("action", "", "create,migrate")
	name := flag.String("name", "", "(create) Name of the migration")

	flag.Parse()

	switch *action {
	case "create":
		createMigrationFile(*name)
	case "migrate":
		core.RunMigrations()
	}

}

func createMigrationFile(name string) {
	path := core.RootPath() + "/database/migrations"
	version := time.Now().UnixNano()

	path = fmt.Sprintf("%v/%v_%v", path, version, name)

	createUp, err := os.Create(path + ".up.sql")
	if err != nil {
		log.Fatalln(err)
		return
	}
	createUp.Close()

	createDown, err := os.Create(path + ".down.sql")
	if err != nil {
		log.Fatalln(err)
		return
	}

	createDown.Close()
}
