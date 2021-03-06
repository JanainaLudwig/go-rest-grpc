package config

import (
	"github.com/newrelic/go-agent/v3/newrelic"
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
	NewRelicLicence string
	NewRelicApp *newrelic.Application
	ServerGrpc string
	ServerRest string
}

func (a *AppConfig) SetNewrelicApp(app *newrelic.Application) {
	a.NewRelicApp = app
}

func LoadEnv(path string) {
	load(path)

	App.AppEnv = loadString("APP_ENV", envStr("development"))
	App.ApiPort = loadString("API_PORT", envStr("8080"))
	App.GrpcPort = loadString("GRPC_PORT", envStr("9000"))
	//App.Debug = App.AppEnv == "development"
	App.Debug = false

	App.DbConfig = &database.Config{
		Host:         loadString("MYSQL_HOST", nil),
		User:         loadString("MYSQL_USER", nil),
		Database:     loadString("MYSQL_DATABASE", nil),
		Port:         loadString("MYSQL_PORT", nil),
		Password:     loadString("MYSQL_PASSWORD", nil),
		MaxOpenConns: 100,
		MaxIdleConns: 100,
		MaxLifetime:  4 * time.Minute,
	}

	App.NewRelicLicence = loadString("NEW_RELIC_LICENCE", nil)
	App.ServerGrpc = loadString("SERVER_GRPC", nil)
	App.ServerRest = loadString("SERVER_REST", nil)
}

func RootPath() string {
	_, file, _, _ := runtime.Caller(0)

	root := path.Dir(path.Dir(file))

	return root
}