package config

import (
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	ApiToken string
}

var Config config

func InitEnvironment() config {
	if err := godotenv.Load(); err != nil {
		panic("error loading .env file")
	}

	if apiToken := os.Getenv("API_TOKEN"); apiToken == "" {
		panic("API_TOKEN environment is required")
	}
	Config.ApiToken = os.Getenv("API_TOKEN")
	return Config
}
