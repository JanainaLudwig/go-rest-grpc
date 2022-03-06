package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type Config struct {
	Host string
	User string
	Database string
	Port string
	Password string
	MaxOpenConns int
	MaxIdleConns int
	MaxLifetime time.Duration
}

func NewConnectionPool(config *Config) *sql.DB {
	db, err := sql.Open("mysql", config.User+":"+config.Password+"@tcp("+config.Host+":"+config.Port+")/"+config.Database+"?parseTime=true")
	if err != nil {
		log.Println("Error opening db connection", err)
		return nil
	}

	db.SetConnMaxLifetime(config.MaxLifetime)
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)

	return db
}

func CloseRows(rows *sql.Rows) {
	err := rows.Close()
	if err != nil {
		log.Println("Error closing rows", err)
	}
}