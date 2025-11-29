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

func TestQuestionUseCase_Create(t *testing.T) {
	// Тихий логгер, чтобы не мусорить в консоль во время тестов
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))

	type args struct {
		ctx context.Context
		q   *domain.Question
	}

	ctxAdmin := domain.NewContextWithUser(context.Background(), domain.ID(uuid.New().String()), domain.RoleAdmin)
	ctxUser := domain.NewContextWithUser(context.Background(), domain.ID(uuid.New().String()), domain.RoleUser)
	ctxUnauthorized := context.Background()

	validQuestion := &domain.Question{
		Text:       "2+2=?",
		Type:       domain.TypeMultiChoice,
		Difficulty: 1,
		Options: []domain.Option{
			{Text: "3", IsCorrect: false},
			{Text: "4", IsCorrect: true},
		},
	}

	unValidQuestion := &domain.Question{
		Text:       "",
		Type:       domain.TypeMultiChoice,
		Difficulty: 1,
		Options: []domain.Option{
			{Text: "3", IsCorrect: false},
			{Text: "4", IsCorrect: true},
		},
	}

	tests := []struct {
		name        string
		args        args
		setupMocks  func(qr *mocks.MockQuestionRepository)
		expectedErr error
		expectCall  bool // Ожидаем ли вызов метода репозитория
	}{
		{
			name: "Successful question creation",
			args: args{
				ctx: ctxAdmin,
				q:   validQuestion,
			},
			setupMocks: func(qr *mocks.MockQuestionRepository) {
				qr.EXPECT().Create(mock.Anything, mock.AnythingOfType("*domain.Question")).Return(nil)
			},
			expectedErr: nil,
			expectCall:  true,
		},
		{
			name: "User without admin role tries to create question",
			args: args{
				ctx: ctxUser,
				q:   validQuestion,
			},
			setupMocks: func(qr *mocks.MockQuestionRepository) {
				// В этом тесте метод репозитория не должен вызываться
			},
			expectedErr: domain.ErrForbidden,
			expectCall:  false,
		},
		{
			name: "Unauthorized user tries to create question",
			args: args{
				ctx: ctxUnauthorized,
				q:   validQuestion,
			},
			setupMocks: func(qr *mocks.MockQuestionRepository) {
				// В этом тесте метод репозитория не должен вызываться
			},
			expectedErr: domain.ErrUnauthorized,
			expectCall:  false,
		},
		{
			name: "Validation fails for question",
			args: args{
				ctx: ctxAdmin,
				q:   unValidQuestion,
			},
			setupMocks: func(qr *mocks.MockQuestionRepository) {
				// В этом тесте метод репозитория не должен вызываться
			},
			expectedErr: domain.ErrEmptyText,
			expectCall:  false,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			// Создаем мок репозитория вопросов
			mockQuestionRepo := mocks.NewMockQuestionRepository(t)
			tt.setupMocks(mockQuestionRepo)

			// Создаем юзкейc с моками
			uc := NewQuestionUseCase(mockQuestionRepo, logger)

			// Вызываем тестируемый метод
			err := uc.CreateQuestion(tt.args.ctx, tt.args.q)

			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr)
			} else {
				assert.NoError(t, err)
			}
			mockQuestionRepo.AssertExpectations(t)
		})
	}
}
