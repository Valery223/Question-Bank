package postgres

import (
	"context"
	"database/sql"

	"github.com/Valery223/Question-Bank/internal/domain"
	"github.com/Valery223/Question-Bank/internal/usecase/ports"
)

// TemplateRepository предоставляет методы для работы с тестовыми шаблонами в базе данных Postgres
type TemplateRepository struct {
	db *sql.DB
}

func NewTemplateRepository(db *sql.DB) *TemplateRepository {
	return &TemplateRepository{db: db}
}

func (r *TemplateRepository) Create(ctx context.Context, tt *domain.TestTemplate) error {
	// Транзакция, так как вставляем в 2 таблицы: test_templates и template_questions
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// Вставляем шаблон теста
	queryT := `INSERT INTO test_templates (id, name, role, purpose) VALUES ($1, $2, $3, $4)`

	_, err = tx.ExecContext(ctx, queryT,
		tt.ID, tt.Name, tt.Role, tt.Purpose,
	)
	if err != nil {
		return err
	}

	// Вставляем вопросы шаблона
	for i, questionID := range tt.QuestionIDs {
		queryQ := `INSERT INTO test_template_questions (template_id, question_id, question_order) VALUES ($1, $2, $3)`
		_, err = tx.ExecContext(ctx, queryQ,
			tt.ID, questionID, i,
		)
		if err != nil {
			return err
		}
	}

	// Успешно
	return nil
}

func (r *TemplateRepository) GetByID(ctx context.Context, id domain.ID) (*domain.TestTemplate, error) {

	tt := &domain.TestTemplate{}

	queryT := `SELECT id, name, role, purpose FROM test_templates WHERE id = $1`

	err := r.db.QueryRowContext(ctx, queryT, id).Scan(
		&tt.ID, &tt.Name, &tt.Role, &tt.Purpose,
	)
	if err != nil {
		return nil, err
	}

	// Получаем вопросы шаблона
	queryQ := `SELECT question_id FROM test_template_questions WHERE template_id = $1 ORDER BY question_order`
	rows, err := r.db.QueryContext(ctx, queryQ, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var questionID string
		if err := rows.Scan(&questionID); err != nil {
			return nil, err
		}
		tt.QuestionIDs = append(tt.QuestionIDs, domain.ID(questionID))
	}

	return tt, nil
}

func (r *TemplateRepository) GetByIDs(ctx context.Context, ids []domain.ID) ([]domain.TestTemplate, error) {
	templates := make([]domain.TestTemplate, 0, len(ids))
	for _, id := range ids {
		tt, err := r.GetByID(ctx, id)
		if err != nil {
			return nil, err
		}
		templates = append(templates, *tt)
	}
	return templates, nil
}

func (r *TemplateRepository) Delete(ctx context.Context, id domain.ID) error {
	query := `DELETE FROM test_templates WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *TemplateRepository) Update(ctx context.Context, tt *domain.TestTemplate) error {
	query := `UPDATE test_templates SET name = $1, role = $2, purpose = $3 WHERE id = $4`
	_, err := r.db.ExecContext(ctx, query,
		tt.Name, tt.Role, tt.Purpose, tt.ID,
	)
	return err
}

func (r *TemplateRepository) Filter(ctx context.Context, filter ports.TemplateFilter) ([]domain.TestTemplate, error) {
	return nil, nil
}
