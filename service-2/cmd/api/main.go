package main

import (
	"log"

	"github.com/nikhilnarayanan623/x-tention-crew/service2/pkg/config"
	"github.com/nikhilnarayanan623/x-tention-crew/service2/pkg/di"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load config: ", err)
	}

	server, err := di.InitializeAPI(cfg)
	if err != nil {
		log.Fatal("failed to initialize api: ", err)
	}

	if err := server.Start(); err != nil {
		log.Fatal("failed to start server", err)
	}
}
