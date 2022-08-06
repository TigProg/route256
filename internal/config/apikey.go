package config

import (
	"os"
)

var ApiKey string

func init() {
	path, _ := os.LookupEnv("TELEGRAM_API_KEY")
	ApiKey = path
}
