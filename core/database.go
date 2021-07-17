package core

import (
	"database/sql"
	"grpc-rest/config"
	"grpc-rest/lib/database"
)

var DB *sql.DB

func StartDb() {
	DB = database.NewConnectionPool(config.App.DbConfig)
}
