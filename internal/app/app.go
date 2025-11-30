package app

import (
	"embed"
	"log/slog"

	httpServer "github.com/Valery223/Question-Bank/internal/delivery/http_server"
	v1 "github.com/Valery223/Question-Bank/internal/delivery/http_server/v1"
	"github.com/Valery223/Question-Bank/internal/parseconfig"
	"github.com/Valery223/Question-Bank/internal/repository/postgres"
	"github.com/Valery223/Question-Bank/internal/usecase"
	db "github.com/Valery223/Question-Bank/pkg/postgres"
	"github.com/gin-gonic/gin"
	"github.com/pressly/goose/v3"
)

// Embedding миграций в бинарник
//
//go:embed migrations/*.sql
var embedMigrations embed.FS // Эта переменная будет содержать реальные файлы внутри бинарника

func NewApp(cfg *parseconfig.Config, logger *slog.Logger) *gin.Engine {

	// 1. Подключение к базе данных
	db, err := db.NewPostgresDB(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Name)
	if err != nil {
		logger.Error("failed to connect to database", "error", err)
		panic(err)
	}

	// 2. Миграции
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		logger.Error("goose dialect error", "error", err)
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		logger.Error("goose up migration error", "error", err)
		panic(err)
	}

	logger.Info("database connected and migrated successfully")

	// 3. Инициализация репозиториев
	questionRepo := postgres.NewQuestionRepository(db)
	templateRepo := postgres.NewTemplateRepository(db)
	sessionRepo := postgres.NewTestSessionRepository(db)

	// 4. Инициализация usecase слоев
	questionUC := usecase.NewQuestionUseCase(questionRepo, logger)
	templateUC := usecase.NewTemplateUseCase(templateRepo, questionRepo, logger)
	sessionUC := usecase.NewSessionUseCase(sessionRepo, templateRepo, questionRepo, logger)

	// 5. Инициализация HTTP сервера и маршрутов
	handler := v1.NewHandler(questionUC, templateUC, sessionUC, logger)
	router := httpServer.NewRouter(handler)

	return router
}
