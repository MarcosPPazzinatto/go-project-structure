package config

import (
	"log"
	"os"
)

type Config struct {
	Port string
}

func Load() Config {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
		log.Println("APP_PORT not set, defaulting to 8080")
	}

	return Config{
		Port: port,
	}
}

