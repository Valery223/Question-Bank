package v1

import (
	"log/slog"

	"github.com/Valery223/Question-Bank/internal/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	questionUC usecase.QuestionUseCase
	templateUC usecase.TemplateUseCase
	sessionUC  usecase.SessionUseCase
	logger     *slog.Logger
}

func NewHandler(qUC usecase.QuestionUseCase, tUC usecase.TemplateUseCase, sUC usecase.SessionUseCase, logger *slog.Logger) *Handler {
	return &Handler{
		questionUC: qUC,
		templateUC: tUC,
		sessionUC:  sUC,
		logger:     logger,
	}
}

// Вспомогательная функция для ошибок
func (h *Handler) errorResponse(c *gin.Context, code int, msg string) {
	h.logger.Error("http response error", "status", code, "msg", msg)
	c.AbortWithStatusJSON(code, gin.H{"error": msg})
}
