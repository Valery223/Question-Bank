package ports

import (
	"context"

	"github.com/Valery223/Question-Bank/internal/domain"
)

// QuestionRepository - порт для работы с вопросами
//
// CRUD и фильтрация
type QuestionRepository interface {
	Create(ctx context.Context, q *domain.Question) error
	GetByID(ctx context.Context, id domain.ID) (*domain.Question, error)
	GetByIDs(ctx context.Context, ids []domain.ID) ([]domain.Question, error)
	Delete(ctx context.Context, id domain.ID) error
	Update(ctx context.Context, q *domain.Question) error
	//  Фильтрация вопросов по разным параметрам
	Filter(ctx context.Context, filter QuestionFilter) ([]domain.Question, error)
}

type QuestionFilter struct {
	Role       *domain.RoleQuestionnaire //  Указатель - если nil, то не фильтруем
	Topic      *string
	Difficulty *domain.Difficulty // можно сделать *domain.DifficultyRange
	Types      []domain.QuestionType
	Limit      int
	Offset     int
}
