package main

import (
	"log"

	"github.com/nikhilnarayanan623/x-tention-crew/user-servcie/pkg/config"
	"github.com/nikhilnarayanan623/x-tention-crew/user-servcie/pkg/di"
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

	server.Start()
}
