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
	m := getMigrationClient()
	err := m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			return
		}
		log.Fatalln(err)
		return
	}
}

func DownMigrations() {
	m := getMigrationClient()
	err := m.Down()
	if err != nil {
		if err == migrate.ErrNoChange {
			return
		}
		log.Fatalln(err)
		return
	}
}

func getMigrationClient() *migrate.Migrate {

	db :=  database.NewConnectionPoolMulti(config.App.DbConfig)
	driver, err := mysql.WithInstance(db, &mysql.Config{
	})
	if err != nil {
		log.Fatalln(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///" + RootPath() + "/database/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalln(err)
	}

	return m
}
