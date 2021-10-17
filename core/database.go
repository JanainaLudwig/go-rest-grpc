package core

import (
	"database/sql"
	"grpc-rest/config"
	"grpc-rest/lib/database"
	"log"
)

var DB *sql.DB

func StartDb() {
	DB = database.NewConnectionPool(config.App.DbConfig)
}

func DbClose(rows *sql.Rows) {
	err := rows.Close()
	if err != nil {
		log.Printf("error closing db rows: %v", err)
	}
}