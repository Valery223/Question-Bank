package usecase

import (
	"context"
	"log/slog"

	"github.com/Valery223/Question-Bank/internal/domain"
	"github.com/Valery223/Question-Bank/internal/usecase/ports"
)

// QuestionUseCase - группирует логику вокруг вопросов
type QuestionUseCase struct {
	questionRepo ports.QuestionRepository
	log          *slog.Logger
}

// NewQuestionUseCase - конструктор для QuestionUseCase
func NewQuestionUseCase(qr ports.QuestionRepository, logger *slog.Logger) *QuestionUseCase {
	return &QuestionUseCase{
		questionRepo: qr,
		log:          logger,
	}
}

// CreateQuestion - создает вопрос
func (uc *QuestionUseCase) CreateQuestion(ctx context.Context, q *domain.Question) error {
	uc.log.Info("Creating question", "question", q)

	// Валидация вопроса
	if err := q.Validate(); err != nil {
		uc.log.Error("Validation failed", "error", err)
		return err
	}

	// Бизнес-логика по созданию вопроса
	// Например, можно ли создавать вопросы  для данной роли
	userID, userRole, ok := domain.UserFromContext(ctx)
	if !ok {
		uc.log.Error("Failed to get user from context")
		return domain.ErrUnauthorized
	}

	if !userRole.CanCreateQuestions() {
		uc.log.Warn("User does not have permission to create questions", "userID", userID)
		return domain.ErrForbidden
	}

	return uc.questionRepo.Create(ctx, q)
}

// GetQuestionByID - получает вопрос по ID
func (uc *QuestionUseCase) GetQuestionByID(ctx context.Context, id domain.ID) (*domain.Question, error) {
	uc.log.Info("Getting question by ID", "id", id)

	userID, userRole, ok := domain.UserFromContext(ctx)
	if !ok {
		uc.log.Error("Failed to get user from context")
		return nil, domain.ErrUnauthorized
	}

	if !userRole.CanViewQuestions() {
		uc.log.Warn("User does not have permission to view questions", "userID", userID)
		return nil, domain.ErrForbidden
	}

	q, err := uc.questionRepo.GetByID(ctx, id)

	if err != nil {
		uc.log.Error("Failed to get question", "error", err)
		return nil, err
	}

	return q, nil

}

// UpdateQuestion - обновляет вопрос
func (uc *QuestionUseCase) UpdateQuestion(ctx context.Context, q *domain.Question) error {
	uc.log.Info("Updating question", "question", q)

	// Валидация вопроса
	if err := q.Validate(); err != nil {
		uc.log.Error("Validation failed", "error", err)
		return err
	}

	// Бизнес-логика по обновлению вопроса
	userID, userRole, ok := domain.UserFromContext(ctx)
	if !ok {
		uc.log.Error("Failed to get user from context")
		return domain.ErrUnauthorized
	}

	if !userRole.CanUpdateQuestions() {
		uc.log.Warn("User does not have permission to update questions", "userID", userID)
		return domain.ErrForbidden
	}

	return uc.questionRepo.Update(ctx, q)
}

func (uc *QuestionUseCase) DeleteQuestion(ctx context.Context, id domain.ID) error {
	uc.log.Info("Deleting question", "id", id)

	// Бизнес-логика по удалению вопроса
	userID, userRole, ok := domain.UserFromContext(ctx)
	if !ok {
		uc.log.Error("Failed to get user from context")
		return domain.ErrUnauthorized
	}

	if !userRole.CanDeleteQuestions() {
		uc.log.Warn("User does not have permission to delete questions", "userID", userID)
		return domain.ErrForbidden
	}

	return uc.questionRepo.Delete(ctx, id)
}
