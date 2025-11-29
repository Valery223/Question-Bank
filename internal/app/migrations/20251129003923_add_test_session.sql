-- +goose Up
CREATE TABLE test_sessions (
    id           UUID PRIMARY KEY,
    template_id  UUID NOT NULL REFERENCES test_templates(id),
    user_id      UUID, 
    started_at   TIMESTAMP NOT NULL,
    expired_at   TIMESTAMP NOT NULL,
    
    -- Мы храним массив вопросов прямо здесь
    -- Это гарантирует, что даже если админ удалит вопрос из базы,
    -- в этой сессии он останется навсегда
    questions_snapshot JSONB NOT NULL
);


-- +goose Down
DROP TABLE test_sessions;