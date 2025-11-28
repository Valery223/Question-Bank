package memory

import (
	"context"
	"errors"

	"github.com/Valery223/Question-Bank/internal/domain"
	"github.com/Valery223/Question-Bank/internal/usecase/ports"
)

type MemoryRepository struct {
	Questions map[domain.ID]domain.Question
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		Questions: make(map[domain.ID]domain.Question),
	}
}

func (r *MemoryRepository) Create(ctx context.Context, q *domain.Question) error {
	r.Questions[q.ID] = *q
	return nil
}

func (r *MemoryRepository) GetByID(ctx context.Context, id domain.ID) (*domain.Question, error) {
	q, exists := r.Questions[id]
	if !exists {
		return nil, errors.New("question not found")
	}
	return &q, nil
}

func (r *MemoryRepository) GetByIDs(ctx context.Context, ids []domain.ID) ([]domain.Question, error) {
	var questions []domain.Question
	for _, id := range ids {
		if q, exists := r.Questions[id]; exists {
			questions = append(questions, q)
		}
	}
	return questions, nil
}

func (r *MemoryRepository) Delete(ctx context.Context, id domain.ID) error {
	delete(r.Questions, id)
	return nil
}

func (r *MemoryRepository) Update(ctx context.Context, q *domain.Question) error {
	r.Questions[q.ID] = *q
	return nil
}

func (r *MemoryRepository) Filter(ctx context.Context, filter ports.QuestionFilter) ([]domain.Question, error) {
	return nil, nil
}
