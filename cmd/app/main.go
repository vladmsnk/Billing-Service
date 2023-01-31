package main

import (
	"log"

	"vladmsnk/billing/config"
	"vladmsnk/billing/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
