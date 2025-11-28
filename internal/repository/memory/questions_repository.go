package memory

import (
	"context"
	"errors"

	"github.com/Valery223/Question-Bank/internal/domain"
	"github.com/Valery223/Question-Bank/internal/usecase/ports"
)

type QuestionsRepository struct {
	storage *MemoryRepository
}

func NewQuestionsRepository(storage *MemoryRepository) *QuestionsRepository {
	return &QuestionsRepository{
		storage: storage,
	}
}

func (r *QuestionsRepository) Create(ctx context.Context, q *domain.Question) error {
	r.storage.Questions[q.ID] = *q
	return nil
}

func (r *QuestionsRepository) GetByID(ctx context.Context, id domain.ID) (*domain.Question, error) {
	q, exists := r.storage.Questions[id]
	if !exists {
		return nil, errors.New("question not found")
	}
	return &q, nil
}

func (r *QuestionsRepository) GetByIDs(ctx context.Context, ids []domain.ID) ([]domain.Question, error) {
	var questions []domain.Question
	for _, id := range ids {
		if q, exists := r.storage.Questions[id]; exists {
			questions = append(questions, q)
		}
	}
	return questions, nil
}

func (r *QuestionsRepository) Delete(ctx context.Context, id domain.ID) error {
	delete(r.storage.Questions, id)
	return nil
}

func (r *QuestionsRepository) Update(ctx context.Context, q *domain.Question) error {
	r.storage.Questions[q.ID] = *q
	return nil
}

func (r *QuestionsRepository) Filter(ctx context.Context, filter ports.QuestionFilter) ([]domain.Question, error) {
	return nil, nil
}
