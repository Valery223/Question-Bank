package memory

import (
	"context"
	"errors"

	"github.com/Valery223/Question-Bank/internal/domain"
	"github.com/Valery223/Question-Bank/internal/usecase/ports"
)

type TemplateRepository struct {
	storage *MemoryRepository
}

func NewTemplateRepository(storage *MemoryRepository) *TemplateRepository {
	return &TemplateRepository{
		storage: storage,
	}
}

func (r *TemplateRepository) Create(ctx context.Context, template *domain.TestTemplate) error {
	r.storage.Templates[template.ID] = *template
	return nil
}

func (r *TemplateRepository) GetByID(ctx context.Context, id domain.ID) (*domain.TestTemplate, error) {
	template, exists := r.storage.Templates[id]
	if !exists {
		return nil, errors.New("template not found")
	}
	return &template, nil
}

func (r *TemplateRepository) GetByIDs(ctx context.Context, ids []domain.ID) ([]domain.TestTemplate, error) {
	var templates []domain.TestTemplate
	for _, id := range ids {
		if template, exists := r.storage.Templates[id]; exists {
			templates = append(templates, template)
		}
	}
	return templates, nil
}

func (r *TemplateRepository) Delete(ctx context.Context, id domain.ID) error {
	delete(r.storage.Templates, id)
	return nil
}
func (r *TemplateRepository) Update(ctx context.Context, template *domain.TestTemplate) error {
	r.storage.Templates[template.ID] = *template
	return nil
}

func (r *TemplateRepository) Filter(ctx context.Context, filter ports.TemplateFilter) ([]domain.TestTemplate, error) {
	return nil, nil
}
