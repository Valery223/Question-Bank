package app

import (
	"log/slog"

	httpServer "github.com/Valery223/Question-Bank/internal/delivery/http_server"
	v1 "github.com/Valery223/Question-Bank/internal/delivery/http_server/v1"
	"github.com/Valery223/Question-Bank/internal/parseconfig"
	"github.com/Valery223/Question-Bank/internal/repository/postgres"
	"github.com/Valery223/Question-Bank/internal/usecase"
	db "github.com/Valery223/Question-Bank/pkg/postgres"
	"github.com/gin-gonic/gin"
)

func NewApp(cfg *parseconfig.Config, logger *slog.Logger) *gin.Engine {

	// Инициализация репозиториев
	// memRepo := memory.NewMemoryRepository()
	db, err := db.NewPostgresDB(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Name)
	if err != nil {
		logger.Error("failed to connect to database", "error", err)
		panic(err)
	}

	// questionRepo := memory.NewQuestionsRepository(memRepo)
	// templateRepo := memory.NewTemplateRepository(memRepo)
	questionRepo := postgres.NewQuestionRepository(db)
	templateRepo := postgres.NewTemplateRepository(db)
	sessionRepo := postgres.NewTestSessionRepository(db)

	// Инициализация usecase слоев
	questionUC := usecase.NewQuestionUseCase(questionRepo, logger)
	templateUC := usecase.NewTemplateUseCase(templateRepo, questionRepo, logger)
	sessionUC := usecase.NewSessionUseCase(sessionRepo, templateRepo, questionRepo, logger)
	// Инициализация HTTP сервера и маршрутов
	handler := v1.NewHandler(questionUC, templateUC, sessionUC, logger)
	router := httpServer.NewRouter(handler)

	return router
}
