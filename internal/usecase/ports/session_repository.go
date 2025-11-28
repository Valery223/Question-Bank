package ports

import (
	"context"

	"github.com/Valery223/Question-Bank/internal/domain"
)

// TestSessionRepository - порт для работы с сессиями(тестирования)
type TestSessionRepository interface {
	CreateSession(ctx context.Context, session *domain.TestSession) error
	GetSession(ctx context.Context, id domain.ID) (*domain.TestSession, error)
	DeleteSession(ctx context.Context, id domain.ID) error
}
