package config

import "github.com/joho/godotenv"

func initEnvs() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}
