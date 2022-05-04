package core

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"grpc-rest/config"
	"grpc-rest/lib/database"
	"log"
)

func RunMigrations() {
	m := getMigrationClient("migrations")
	if m == nil {
		log.Println("failed running migrations")
		return
	}
	err := m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			return
		}
		log.Println(err)
		return
	}
}

func DownMigrations() {
	m := getMigrationClient("migrations")
	err := m.Down()
	if err != nil {
		if err == migrate.ErrNoChange {
			return
		}
		log.Println(err)
		return
	}
}

func getMigrationClient(path string) *migrate.Migrate {

	db :=  database.NewConnectionPoolMulti(config.App.DbConfig)
	driver, err := mysql.WithInstance(db, &mysql.Config{
	})
	if err != nil {
		log.Println(err)
		return nil
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///" + RootPath() + "/database/" + path,
		"postgres", driver)
	if err != nil {
		log.Println(err)
	}

	return m
}
