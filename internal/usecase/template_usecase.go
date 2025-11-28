package usecase

import (
	"context"
	"errors"
	"log/slog"

	"github.com/Valery223/Question-Bank/internal/domain"
	"github.com/Valery223/Question-Bank/internal/usecase/ports"
	"github.com/google/uuid"
)

type TemplateUseCase struct {
	templateRepo ports.TemplateRepository
	questionRepo ports.QuestionRepository
	log          *slog.Logger
}

func NewTemplateUseCase(tr ports.TemplateRepository, qr ports.QuestionRepository, logger *slog.Logger) *TemplateUseCase {
	return &TemplateUseCase{
		templateRepo: tr,
		questionRepo: qr,
		log:          logger,
	}
}

// CreateTemplate - создает шаблон теста
//
// Проверяет, что все вопросы в шаблоне существуют
// Можно отдельно создать usecase который будет принимать TestTemplate с новыми вопросами
func (uc *TemplateUseCase) CreateTemplate(ctx context.Context, tt *domain.TestTemplate) error {
	uc.log.Info("Creating test template", "template", tt)

	// Валидация шаблона
	if err := tt.Validate(); err != nil {
		uc.log.Error("Validation failed", "error", err)
		return err
	}

	// Бизнес-логика по созданию шаблона
	// Например, можно ли создавать шаблоны для данной роли
	userID, userRole, ok := domain.UserFromContext(ctx)
	if !ok {
		uc.log.Error("Failed to get user from context")
		return domain.ErrUnauthorized
	}

	if !userRole.CanCreateTemplates() {
		uc.log.Warn("User does not have permission to create templates", "userID", userID)
		return domain.ErrForbidden
	}

	// Проверяем, что все вопросы существуют
	questions, err := uc.questionRepo.GetByIDs(ctx, tt.QuestionIDs)
	if err != nil {
		uc.log.Error("Failed to get questions for template", "error", err)
		return err
	}
	if len(questions) != len(tt.QuestionIDs) {
		uc.log.Warn("Some questions do not exist for the template", "expected", len(tt.QuestionIDs), "found", len(questions))
		return errors.New("some questions do not exist for the template")
	}

	id := uuid.New().String()
	tt.ID = domain.ID(id)
	return uc.templateRepo.Create(ctx, tt)
}

func (uc *TemplateUseCase) GetTemplateByID(ctx context.Context, id domain.ID) (*domain.TestTemplate, error) {
	uc.log.Info("Getting test template by ID", "id", id)

	userID, userRole, ok := domain.UserFromContext(ctx)
	if !ok {
		uc.log.Error("Failed to get user from context")
		return nil, domain.ErrUnauthorized
	}

	if !userRole.CanViewTemplates() {
		uc.log.Warn("User does not have permission to view templates", "userID", userID)
		return nil, domain.ErrForbidden
	}

	return uc.templateRepo.GetByID(ctx, id)
}

func (uc *TemplateUseCase) GetTemplateDetailsByID(ctx context.Context, id domain.ID) (*domain.TestTemplate, []domain.Question, error) {
	uc.log.Info("Getting test template details by ID", "id", id)

	userID, userRole, ok := domain.UserFromContext(ctx)
	if !ok {
		uc.log.Error("Failed to get user from context")
		return nil, nil, domain.ErrUnauthorized
	}

	if !userRole.CanViewTemplates() {
		uc.log.Warn("User does not have permission to view templates", "userID", userID)
		return nil, nil, domain.ErrForbidden
	}
	template, err := uc.templateRepo.GetByID(ctx, id)
	if err != nil {
		uc.log.Error("Failed to get template", "error", err)
		return nil, nil, err
	}

	questions, err := uc.questionRepo.GetByIDs(ctx, template.QuestionIDs)
	if err != nil {
		uc.log.Error("Failed to get questions for template", "error", err)
		return nil, nil, err
	}

	return template, questions, nil
}

func (uc *TemplateUseCase) DeleteTemplate(ctx context.Context, id domain.ID) error {
	uc.log.Info("Deleting test template", "id", id)

	userID, userRole, ok := domain.UserFromContext(ctx)
	if !ok {
		uc.log.Error("Failed to get user from context")
		return domain.ErrUnauthorized
	}

	if !userRole.CanDeleteTemplates() {
		uc.log.Warn("User does not have permission to delete templates", "userID", userID)
		return domain.ErrForbidden
	}

	return uc.templateRepo.Delete(ctx, id)
}
