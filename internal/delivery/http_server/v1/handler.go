package v1

import (
	"log/slog"

	"github.com/Valery223/Question-Bank/internal/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	questionUC usecase.QuestionUseCase
	logger     *slog.Logger
}

func NewHandler(qUC usecase.QuestionUseCase, logger *slog.Logger) *Handler {
	return &Handler{
		questionUC: qUC,
		logger:     logger,
	}
}

// Вспомогательная функция для ошибок
func (h *Handler) errorResponse(c *gin.Context, code int, msg string) {
	h.logger.Error("http response error", "status", code, "msg", msg)
	c.AbortWithStatusJSON(code, gin.H{"error": msg})
}
