package main

import (
	"context"
	"flag"

	_ "github.com/lib/pq"
	"github.com/nevano11/minimarket/internal/minimarket/bootstrap"
	log "github.com/sirupsen/logrus"
)

func main() {
	var (
		configPath string
		logLevel   string
	)

	flag.StringVar(&configPath, "config", "cmd/minimarket/config-example.yaml", "path to config")
	flag.StringVar(&logLevel, "log-level", "debug", "level of logging")
	flag.Parse()

	level, err := log.ParseLevel(logLevel)
	if err != nil {
		log.Error("failed to parse log-level")
		return
	}
	log.SetLevel(level)

	ctx := context.Background()

	err = bootstrap.Run(ctx, bootstrap.Config{
		ConfigPath: configPath,
		LogLevel:   logLevel,
	})

	if err != nil {
		log.Error(err)
	}

	log.Info("minimarket stopped")
}
