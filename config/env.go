package config

import (
	"os"

	"github.com/joho/godotenv"
)

var env *Environments

type Environments struct {
	APP_PORT string
	URI_DB   string
	DATABASE string
}

func LoadEnvironments() *Environments {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	env = &Environments{
		APP_PORT: os.Getenv("APP_PORT"),
		URI_DB:   os.Getenv("URI_DB"),
		DATABASE: os.Getenv("DATABASE"),
	}

	return env
}
