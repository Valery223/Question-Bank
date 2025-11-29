package postgres_test

import (
	"context"
	"testing"

	"github.com/Valery223/Question-Bank/internal/domain"
	"github.com/Valery223/Question-Bank/internal/repository/postgres"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQuestionRepo_CRUD(t *testing.T) {
	db := SetupTestDB(t)
	defer db.Close()

	repo := postgres.NewQuestionRepository(db)
	ctx := context.Background()

	// 2. Данные для теста
	q := &domain.Question{
		ID:         domain.ID(uuid.New().String()),
		Text:       "Integration Test Question",
		Role:       "backend_junior",
		Topic:      "golang",
		Type:       domain.TypeSingleChoice,
		Difficulty: 3,
		Options: []domain.Option{
			{ID: "opt-1", Text: "Yes", IsCorrect: true},
			{ID: "opt-2", Text: "No", IsCorrect: false},
		},
	}

	// --- TEST CREATE ---
	t.Run("Create Question", func(t *testing.T) {
		err := repo.Create(ctx, q)
		require.NoError(t, err)
	})

	// --- TEST GET ---
	t.Run("Get Question", func(t *testing.T) {
		fetchedQ, err := repo.GetByID(ctx, q.ID)
		require.NoError(t, err)

		// Проверяем поля
		assert.Equal(t, q.Text, fetchedQ.Text)
		assert.Equal(t, q.Difficulty, fetchedQ.Difficulty)

		// Проверяем, что опции тоже сохранились (связная таблица)
		assert.Len(t, fetchedQ.Options, 2)
		assert.Equal(t, "Yes", fetchedQ.Options[0].Text)
	})
}
