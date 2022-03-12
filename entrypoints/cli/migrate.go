package main

import (
	"flag"
	"fmt"
	"grpc-rest/core"
	"log"
	"os"
	"time"
)

func main()  {
	action := flag.String("action", "", "create")
	 name := flag.String("name", "", "Name of the migration")

	flag.Parse()

	switch action {
	case action:
		createMigrationFile(*name)
	}

}

func createMigrationFile(name string) {
	path := core.RootPath() + "/database/migrations"
	up := time.Now().UnixNano()

	path = fmt.Sprintf("%v/%v_%v", path, up, name)

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
