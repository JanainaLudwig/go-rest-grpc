package core

import (
	"path/filepath"
	"runtime"
)

func StartApp(runMigration bool) {
	//StartDb()
	//if runMigration {
	//	RunMigrations()
	//}
}

func RootPath() string {
	_, b, _, _ := runtime.Caller(0)
	core := filepath.Dir(b)
	return filepath.Dir(core)
}
