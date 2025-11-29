-- +goose Up
-- +goose StatementBegin
CREATE TABLE test_templates (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    role VARCHAR(50) NOT NULL,
    purpose VARCHAR(50) NOT NULL
);

-- Таблица связи (Many-to-Many)
CREATE TABLE test_template_questions (
    template_id     UUID NOT NULL REFERENCES test_templates(id) ON DELETE CASCADE,
    question_id     UUID NOT NULL REFERENCES questions(id) ON DELETE RESTRICT,
    question_order  INT NOT NULL,
    PRIMARY KEY (template_id, question_id)
);

-- +goose Down
-- SQL в этом разделе выполняется при откате

DROP TABLE test_template_questions;
DROP TABLE test_templates;