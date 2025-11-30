package v1

import (
	"context"
	"log/slog"

	"github.com/Valery223/Question-Bank/internal/domain"
	"github.com/gin-gonic/gin"
)

type QuestionUseCase interface {
	CreateQuestion(ctx context.Context, q *domain.Question) error
	GetQuestionByID(ctx context.Context, id domain.ID) (*domain.Question, error)
	UpdateQuestion(ctx context.Context, q *domain.Question) error
	DeleteQuestion(ctx context.Context, id domain.ID) error
}

type TemplateUseCase interface {
	CreateTemplate(ctx context.Context, t *domain.TestTemplate) error
	GetTemplateByID(ctx context.Context, id domain.ID) (*domain.TestTemplate, error)
	UpdateTemplate(ctx context.Context, t *domain.TestTemplate) error
	DeleteTemplate(ctx context.Context, id domain.ID) error
}

type SessionUseCase interface {
	CreateSession(ctx context.Context, s *domain.TestSession) error
	GetSessionByID(ctx context.Context, id domain.ID) (*domain.TestSession, error)
}

type Handler struct {
	questionUC QuestionUseCase
	templateUC TemplateUseCase
	sessionUC  SessionUseCase
	logger     *slog.Logger
}

func NewHandler(qUC QuestionUseCase, tUC TemplateUseCase, sUC SessionUseCase, logger *slog.Logger) *Handler {
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
