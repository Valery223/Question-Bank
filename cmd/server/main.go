package main

import (
	"github.com/Valery223/Question-Bank/internal/app"
	"github.com/Valery223/Question-Bank/internal/parseconfig"
	"github.com/Valery223/Question-Bank/pkg/logger"
)

// @title           Question Bank API
// @version         1.0
// @description     API сервис для управления банком вопросов и генерации тестов.

// @host            localhost:8080
// @BasePath        /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Загрузка конфигурации
	cfg := parseconfig.MustLoad()

	// Инициализация логгера
	logger := logger.Setup(cfg.Env)

	logger.Info("starting server", "address", cfg.Address, "env", cfg.Env)

	// Инициализация и запуск приложения
	app := app.NewApp(cfg, logger)
	app.Run(cfg.Address)
}
