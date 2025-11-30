// Package memory предоставляет реализацию репозитория в памяти для хранения вопросов и тестовых шаблонов
// Просто для демонстрационных и тестовых целей
package memory

import (
	"github.com/Valery223/Question-Bank/internal/domain"
)

type MemoryRepository struct {
	Questions map[domain.ID]domain.Question
	Templates map[domain.ID]domain.TestTemplate
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		Questions: make(map[domain.ID]domain.Question),
		Templates: make(map[domain.ID]domain.TestTemplate),
	}
}
