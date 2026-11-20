package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct{ // our configurations
	Port string
	Env string
}

func MustLoad() Config { // Mustload means it should either work or kill the program
	godotenv.Load()

	port:= os.Getenv("PORT")
	if port == "" {
		panic("PORT is required")
	}

	env:= os.Getenv("ENV")
	if env == "" {
		panic("ENV is required")
	}

	return Config{
		Port: port,
		Env: env,
	}
}