package postgres

import (
	"context"
	"database/sql"

	"github.com/Valery223/Question-Bank/internal/domain"
	"github.com/Valery223/Question-Bank/internal/usecase/ports"
	"github.com/google/uuid"
)

type QuestionRepository struct {
	db *sql.DB
}

func NewQuestionRepository(db *sql.DB) *QuestionRepository {
	return &QuestionRepository{db: db}
}

func (r *QuestionRepository) Create(ctx context.Context, q *domain.Question) error {
	// Начинаем транзакцию (так как вставляем в 2 таблицы: questions и options)
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

	// Вставляем вопрос
	queryQ := `INSERT INTO questions (id, text, role, topic, type, difficulty) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err = tx.ExecContext(ctx, queryQ,
		q.ID, q.Text, q.Role, q.Topic, q.Type, q.Difficulty,
	)
	if err != nil {
		return err
	}

	// Вставляем варианты ответа (если есть)
	for _, option := range q.Options {
		queryO := `INSERT INTO options (id, question_id, text, is_correct) VALUES ($1, $2, $3, $4)`
		// Генерируем новый UUID для опции
		//
		// Либо можно генерировать в самой БД с помощью DEFAULT gen_random_uuid()
		idOption := uuid.New().String()

		_, err = tx.ExecContext(ctx, queryO,
			idOption, q.ID, option.Text, option.IsCorrect,
		)
		if err != nil {
			return err
		}
	}

	// Успешно
	return nil
}

func (r *QuestionRepository) GetByID(ctx context.Context, id domain.ID) (*domain.Question, error) {
	q := &domain.Question{}
	queryQ := `SELECT id, text, role, topic, type, difficulty FROM questions WHERE id = $1`

	err := r.db.QueryRowContext(ctx, queryQ, id).Scan(
		&q.ID, &q.Text, &q.Role, &q.Topic, &q.Type, &q.Difficulty,
	)
	if err != nil {
		return nil, err
	}

	// Получаем варианты ответа
	queryO := `SELECT id, text, is_correct FROM options WHERE question_id = $1`
	rows, err := r.db.QueryContext(ctx, queryO, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var option domain.Option

		if err := rows.Scan(&option.ID, &option.Text, &option.IsCorrect); err != nil {
			return nil, err
		}
		q.Options = append(q.Options, option)
	}

	return q, nil
}

func (r *QuestionRepository) GetByIDs(ctx context.Context, ids []domain.ID) ([]domain.Question, error) {
	questions := make([]domain.Question, 0, len(ids))
	for _, id := range ids {
		q, err := r.GetByID(ctx, id)
		if err != nil {
			return nil, err
		}
		questions = append(questions, *q)
	}
	return questions, nil
}
func (r *QuestionRepository) Delete(ctx context.Context, id domain.ID) error {
	return nil
}
func (r *QuestionRepository) Update(ctx context.Context, q *domain.Question) error {
	return nil
}
func (r *QuestionRepository) Filter(ctx context.Context, filter ports.QuestionFilter) ([]domain.Question, error) {
	return nil, nil
}
