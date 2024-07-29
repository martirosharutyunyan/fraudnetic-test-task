package config

import (
	"github.com/subosito/gotenv"
	"os"
	"strconv"
)

func Load(files ...string) error {
	return gotenv.Load(files...)
}

func toInt(str string) (int, error) {
	return strconv.Atoi(str)
}

func GetEnv() string {
	return os.Getenv("ENVIRONMENT")
}

func GetDBHost() string {
	host := os.Getenv("DB_HOST")

	return host
}
