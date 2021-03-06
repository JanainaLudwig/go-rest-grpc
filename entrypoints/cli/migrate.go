package main

import (
	"context"
	"flag"
	"fmt"
	"grpc-rest/config"
	"grpc-rest/core"
	"grpc-rest/database/seed"
	"log"
	"os"
	"time"
)

func main() {
	config.LoadEnv(config.RootPath() + "/config/.env")
	core.StartApp(true)
	ctx := context.Background()

	action := flag.String("action", "", "create,migrate,migrate:down,seed")
	name := flag.String("name", "", "(create) Name of the migration")

	flag.Parse()

	switch *action {
	case "create":
		createMigrationFile(*name)
	case "migrate":
		core.RunMigrations()
	case "seed":
		seed.RunSeed(ctx)
	case "migrate:down":
		core.DownMigrations()
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
