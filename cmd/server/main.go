package main

import (
	"github.com/Valery223/Question-Bank/internal/app"
	"github.com/Valery223/Question-Bank/internal/parseconfig"
	"github.com/Valery223/Question-Bank/pkg/logger"
)

func main() {
	cfg := parseconfig.MustLoad()

	logger := logger.Setup(cfg.Env)

	logger.Info("starting server", "address", cfg.Address, "env", cfg.Env)
	app := app.NewApp(cfg, logger)
	app.Run(cfg.Address)
}
