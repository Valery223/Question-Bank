package usecase

import (
	"context"
	"io"
	"log/slog"
	"testing"

	"github.com/Valery223/Question-Bank/internal/domain"
	"github.com/Valery223/Question-Bank/internal/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTemplateUseCase_Create(t *testing.T) {
	// Тихий логгер, чтобы не мусорить в консоль во время тестов
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))

	type args struct {
		ctx context.Context
		q   *domain.TestTemplate
	}

	ctxAdmin := domain.NewContextWithUser(context.Background(), domain.ID(uuid.New().String()), domain.RoleAdmin)
	ctxUser := domain.NewContextWithUser(context.Background(), domain.ID(uuid.New().String()), domain.RoleUser)
	// ctxUnauthorized := context.Background()

	validQuestion := &domain.Question{
		ID:         domain.ID(uuid.New().String()),
		Text:       "2+2=?",
		Type:       domain.TypeMultiChoice,
		Difficulty: 1,
		Options: []domain.Option{
			{Text: "3", IsCorrect: false},
			{Text: "4", IsCorrect: true},
		},
	}

	validTemplate := &domain.TestTemplate{
		Name:        "Sample Test",
		Role:        domain.RoleBackendJunior,
		Purpose:     domain.PurposeMockInterview,
		QuestionIDs: []domain.ID{validQuestion.ID},
	}

	// unValidTemplate := &domain.TestTemplate{
	// 	Name:        "",
	// 	Role:        domain.RoleBackendJunior,
	// 	Purpose:     domain.PurposeMockInterview,
	// 	QuestionIDs: []domain.ID{domain.ID(uuid.New().String()), domain.ID(uuid.New().String())},
	// }

	validMockFunc := func(tr *mocks.MockTemplateRepository, qr *mocks.MockQuestionRepository) {
		tr.EXPECT().Create(mock.Anything, mock.Anything).Return(nil).Once()
		qr.EXPECT().GetByIDs(mock.Anything, validTemplate.QuestionIDs).Return([]domain.Question{
			*validQuestion,
		}, nil).Once()
	}

	tests := []struct {
		name        string
		args        args
		setupMocks  func(tr *mocks.MockTemplateRepository, qr *mocks.MockQuestionRepository)
		expectedErr error
		expectCall  bool // Ожидаем ли вызов метода репозитория
	}{
		{
			name: "Successful template creation",
			args: args{
				ctx: ctxAdmin,
				q:   validTemplate,
			},
			setupMocks:  validMockFunc,
			expectedErr: nil,
			expectCall:  true,
		},
		{
			name: "User without admin role tries to create template",
			args: args{
				ctx: ctxUser,
				q:   validTemplate,
			},
			setupMocks: func(tr *mocks.MockTemplateRepository, qr *mocks.MockQuestionRepository) {
				// В этом тесте методы репозитория не должны вызываться
			},
			expectedErr: domain.ErrForbidden,
			expectCall:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаем мок репозиториев
			mockTemplateRepo := new(mocks.MockTemplateRepository)
			mockQuestionRepo := new(mocks.MockQuestionRepository)

			// Настраиваем моки согласно сценарию теста
			tt.setupMocks(mockTemplateRepo, mockQuestionRepo)

			// Создаем usecase с моками
			uc := NewTemplateUseCase(mockTemplateRepo, mockQuestionRepo, logger)
			err := uc.CreateTemplate(tt.args.ctx, tt.args.q)

			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr)
			} else {
				assert.NoError(t, err)
			}

			// Проверяем, были ли вызваны методы репозитория, если это ожидается
			if tt.expectCall {
				mockTemplateRepo.AssertExpectations(t)
				mockQuestionRepo.AssertExpectations(t)
			}
		})
	}
}
