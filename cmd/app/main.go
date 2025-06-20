package main

import (
	"log"

	"internal/config"
	"internal/server"
)

func main() {
	cfg := config.Load()

	s := server.New(cfg)

	if err := s.Start(); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}

