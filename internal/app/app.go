package app

import (
	"log/slog"
	"os"

	httpServer "github.com/Valery223/Question-Bank/internal/delivery/http_server"
	v1 "github.com/Valery223/Question-Bank/internal/delivery/http_server/v1"
	"github.com/Valery223/Question-Bank/internal/repository/memory"
	"github.com/Valery223/Question-Bank/internal/usecase"
	"github.com/gin-gonic/gin"
)

func NewApp() *gin.Engine {

	// Инициализация логгера
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	// Инициализация репозиториев
	repo := memory.NewMemoryRepository()
	questionRepo := memory.NewQuestionsRepository(repo)
	templateRepo := memory.NewTemplateRepository(repo)
	// Инициализация usecase слоев
	questionUC := usecase.NewQuestionUseCase(questionRepo, logger)
	templateUC := usecase.NewTemplateUseCase(templateRepo, questionRepo, logger)

	handler := v1.NewHandler(*questionUC, *templateUC, logger)
	router := httpServer.NewRouter(handler)

	return router
}
