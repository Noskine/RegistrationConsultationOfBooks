package config

import (
	"os"

	"github.com/joho/godotenv"
)

var env *Environments

type Environments struct {
	APP_PORT string
	URI_DB   string
}

func LoadEnvironments() *Environments {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	env = &Environments{
		APP_PORT: os.Getenv("APP_PORT"),
		URI_DB:   os.Getenv("URI_DB"),
	}

	return env
}
