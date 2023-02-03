package main

import (
	"log"

	"github.com/sanchayata-jain/docker-image/internal/config"
	"github.com/sanchayata-jain/docker-image/internal/server"
)

func main() {
	conf := mustLoadConfig()
	if err := server.New(conf); err != nil {
		log.Fatal(err)
	}
}

func mustLoadConfig() *config.Config {
	conf, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	return conf
}
