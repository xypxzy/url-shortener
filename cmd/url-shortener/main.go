package main

import (
	"fmt"
	"os"
	"url-shortener/internal/config"

	"log/slog"
)

const (
	envLocal      = "local"
	envStaging    = "staging"
	envProduction = "production"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// TODO: init logger - slog
	log := setupLogger(cfg.Env)
	log.Info("Starting the application", slog.String("env", cfg.Env))
	log.Debug("debug message available only in local environment")

	// TODO: init storage - sqllite

	// TODO: init router - chi, render

	// TODO: start server - http

}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envLocal:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envStaging:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProduction:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return logger
}
