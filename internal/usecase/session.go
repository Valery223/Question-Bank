package usecase

import (
	"context"
	"log/slog"
	"time"

	"github.com/Valery223/Question-Bank/internal/domain"
	"github.com/Valery223/Question-Bank/internal/usecase/ports"
	"github.com/google/uuid"
)

type SessionUseCase struct {
	sessionRepo  ports.TestSessionRepository
	templateRepo ports.TemplateRepository
	questionRepo ports.QuestionRepository
	log          *slog.Logger
}

func NewSessionUseCase(sr ports.TestSessionRepository, tr ports.TemplateRepository, qr ports.QuestionRepository, logger *slog.Logger) *SessionUseCase {
	return &SessionUseCase{
		sessionRepo:  sr,
		templateRepo: tr,
		questionRepo: qr,
		log:          logger,
	}
}

func (uc *SessionUseCase) CreateSession(ctx context.Context, session *domain.TestSession) error {
	uc.log.Info("Creating test session")

	// Бизнес-логика по созданию сессии
	userID, userRole, ok := domain.UserFromContext(ctx)
	if !ok {
		uc.log.Error("Failed to get user from context")
		return domain.ErrUnauthorized
	}

	if !userRole.CanCreateSessions() {
		uc.log.Warn("User does not have permission to create sessions", "userID", userID)
		return domain.ErrForbidden
	}

	// Fetch the template
	template, err := uc.templateRepo.GetByID(ctx, session.TemplateID)
	if err != nil {
		uc.log.Error("Failed to fetch template", "error", err)
		return err
	}

	// Create the session
	id := uuid.New().String()
	session.ID = domain.ID(id)
	session.TemplateID = template.ID
	session.UserID = userID
	session.StartedAt = time.Now()
	session.ExpiredAt = time.Now().Add(30 * time.Minute) // Например, сессия длится 30 минут
	session.Questions = make([]domain.Question, len(template.QuestionIDs))

	// Заполняем вопросы из шаблона
	// Так как вопросы могут меняться, копируем их в сессию(делаем снимок)
	questions, err := uc.questionRepo.GetByIDs(ctx, template.QuestionIDs)
	if err != nil {
		uc.log.Error("Failed to fetch questions for session", "error", err)
		return err
	}
	session.Questions = questions

	if err := uc.sessionRepo.CreateSession(ctx, session); err != nil {
		uc.log.Error("Failed to create session", "error", err)
		return err
	}

	uc.log.Info("Test session created successfully", "session_id", session.ID)
	return nil
}

func (uc *SessionUseCase) GetSessionByID(ctx context.Context, id domain.ID) (*domain.TestSession, error) {
	uc.log.Info("Getting test session by ID", "id", id)

	userID, userRole, ok := domain.UserFromContext(ctx)
	if !ok {
		uc.log.Error("Failed to get user from context")
		return nil, domain.ErrUnauthorized
	}

	if !userRole.CanViewSessions() {
		uc.log.Warn("User does not have permission to view sessions", "userID", userID)
		return nil, domain.ErrForbidden
	}

	sessin, err := uc.sessionRepo.GetSession(ctx, id)
	if err != nil {
		uc.log.Error("Failed to get session", "error", err)
		return nil, err
	}

	// Дополнительная проверка: если сессия принадлежит другому пользователю и роль не позволяет просматривать все сессии
	if sessin.UserID != userID && !userRole.CanViewAllSessions() {
		uc.log.Warn("User does not have permission to view this session", "userID", userID, "sessionUserID", sessin.UserID)
		return nil, domain.ErrForbidden
	}

	return sessin, nil
}
