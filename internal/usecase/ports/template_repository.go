package ports

import (
	"context"

	"github.com/Valery223/Question-Bank/internal/domain"
)

// TemplateRepository - порт для работы с шаблонами тестов
//
// CRUD и фильтрация
type TemplateRepository interface {
	Create(ctx context.Context, tt *domain.TestTemplate) error
	GetByID(ctx context.Context, id domain.ID) (*domain.TestTemplate, error)
	GetByIDs(ctx context.Context, ids []domain.ID) ([]domain.TestTemplate, error)
	Delete(ctx context.Context, id domain.ID) error
	Update(ctx context.Context, tt *domain.TestTemplate) error
	//  Фильтрация шаблонов по разным параметрам
	Filter(ctx context.Context, filter TemplateFilter) ([]domain.TestTemplate, error)
}

type TemplateFilter struct {
	Role    *domain.RoleQuestionnaire //  Указатель - если nil, то не фильтруем
	Purpose *domain.TemplatePurpose
	Limit   int
	Offset  int
}
