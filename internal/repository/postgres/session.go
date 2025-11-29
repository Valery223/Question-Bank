package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/Valery223/Question-Bank/internal/domain"
)

type TestSessionRepository struct {
	db *sql.DB
}

// NewTestSessionRepository создает новый репозиторий сессий тестирования
func NewTestSessionRepository(db *sql.DB) *TestSessionRepository {
	return &TestSessionRepository{db: db}
}

func (r *TestSessionRepository) CreateSession(ctx context.Context, session *domain.TestSession) error {
	// Сериализуем вопросы в JSON
	questionsBytes, err := json.Marshal(session.Questions)
	if err != nil {
		return fmt.Errorf("failed to marshal session questions: %w", err)
	}

	query := `
		INSERT INTO test_sessions (id, template_id, user_id, started_at, expired_at, questions_snapshot)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	// ---------
	// Не понятно пока, нужен ли user_id в сессии, поэтому добавил обработку NULL
	// Не ясно, будет ли он или нет, но пока костыль
	var sqlUserID sql.NullString
	if session.UserID != "" {
		sqlUserID = sql.NullString{
			String: string(session.UserID),
			Valid:  true, //  Valid = true для не-NULL значений
		}
	} else {
		sqlUserID = sql.NullString{
			Valid: false, //  Valid = false для NULL
		}
	}

	// ---------

	_, err = r.db.ExecContext(ctx, query,
		session.ID,
		session.TemplateID,
		sqlUserID,
		session.StartedAt,
		session.ExpiredAt,
		questionsBytes,
	)
	if err != nil {
		return fmt.Errorf("failed to create test session: %w", err)
	}

	return nil
}

func (r *TestSessionRepository) GetSession(ctx context.Context, id domain.ID) (*domain.TestSession, error) {
	session := &domain.TestSession{}

	var questionsBytes []byte

	query := `
		SELECT id, template_id, user_id, started_at, expired_at, questions_snapshot
		FROM test_sessions
		WHERE id = $1
	`

	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&session.ID,
		&session.TemplateID,
		&session.UserID,
		&session.StartedAt,
		&session.ExpiredAt,
		&questionsBytes,
	)
	if err != nil {
		// if err == sql.ErrNoRows {
		// 	return nil, domain.ErrSessionNotFound
		// }
		return nil, fmt.Errorf("failed to get test session: %w", err)
	}

	// Десериализуем вопросы из JSON
	err = json.Unmarshal(questionsBytes, &session.Questions)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal session questions: %w", err)
	}

	return session, nil
}

func (r *TestSessionRepository) DeleteSession(ctx context.Context, id domain.ID) error {
	query := `DELETE FROM test_sessions WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
