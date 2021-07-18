package config

import (
	"grpc-rest/lib/database"
	"path"
	"runtime"
	"time"
)

var App AppConfig

type AppConfig struct {
	AppEnv string
	Debug bool
	ApiPort string
	GrpcPort string
	DbConfig *database.Config
}

func LoadEnv(path string) {
	load(path)

	App.AppEnv = loadString("APP_ENV", envStr("development"))
	App.ApiPort = loadString("API_PORT", envStr("8080"))
	App.GrpcPort = loadString("GRPC_PORT", envStr("9000"))
	App.Debug = App.AppEnv == "development"

	App.DbConfig = &database.Config{
		Host:         loadString("MYSQL_HOST", nil),
		User:         loadString("MYSQL_USER", nil),
		Database:     loadString("MYSQL_DATABASE", nil),
		Port:         loadString("MYSQL_PORT", nil),
		Password:     loadString("MYSQL_PASSWORD", nil),
		MaxOpenConns: 20,
		MaxIdleConns: 20,
		MaxLifetime:  4 * time.Minute,
	}
}

func RootPath() string {
	_, file, _, _ := runtime.Caller(0)

	root := path.Dir(path.Dir(file))

	return root
}