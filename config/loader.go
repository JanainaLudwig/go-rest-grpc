package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func load(path string) {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal("Error loading .env from ", path)
	}
}

func envStr(val string) *string {
	return &val
}

func envInt(val int) *int {
	return &val
}

func loadString(code string, defaultValue *string) string {
	val := os.Getenv(code)

	if val == "" {
		return *defaultValue
	}

	return val
}

func loadInt(code string, defaultValue *int) int {
	str := os.Getenv(code)

	if str == "" && defaultValue != nil {
		return *defaultValue
	}

	val, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal("Unable to convert", code, "to int")
	}

	return val
}
