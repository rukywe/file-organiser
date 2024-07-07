package main

import (
	"flag"
	"log"

	"file-organiser/internal/config"
	"file-organiser/internal/organiser"
	"file-organiser/pkg/logger"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	logger := logger.New(cfg.LogLevel)

	org := organiser.New(cfg, logger)
	if err := org.Organize(); err != nil {
		logger.Error("Failed to organize files", "error", err)
	}
}