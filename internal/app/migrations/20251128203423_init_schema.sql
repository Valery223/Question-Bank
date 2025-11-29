-- +goose Up

-- 1. Таблица Вопросов
CREATE TABLE questions (
    id           UUID PRIMARY KEY, -- Мы не используем SERIAL, мы ждем UUID от приложения
    text         TEXT NOT NULL,
    role         VARCHAR(50) NOT NULL,
    topic        VARCHAR(100) NOT NULL,
    type         VARCHAR(20) NOT NULL,
    difficulty   INT NOT NULL
);

-- 2. Таблица Опций (для вопросов с выбором)
CREATE TABLE options (
    id           UUID PRIMARY KEY,
    question_id  UUID NOT NULL REFERENCES questions(id) ON DELETE CASCADE,
    text         TEXT NOT NULL,
    is_correct   BOOLEAN NOT NULL DEFAULT false
);

-- 3. Таблица Шаблонов (Сама "шапка")
CREATE TABLE test_templates (
    id       UUID PRIMARY KEY,
    name     TEXT NOT NULL,
    role     VARCHAR(50) NOT NULL,
    purpose  VARCHAR(50) NOT NULL
);

-- 4. Таблица Связи (Many-to-Many)
CREATE TABLE test_template_questions (
    template_id     UUID NOT NULL REFERENCES test_templates(id) ON DELETE CASCADE,
    question_id     UUID NOT NULL REFERENCES questions(id) ON DELETE RESTRICT,
    question_order  INT NOT NULL,

    -- Составной первичный ключ
    PRIMARY KEY (template_id, question_id)
);

-- +goose Down
DROP TABLE test_template_questions;
DROP TABLE test_templates;
DROP TABLE options;
DROP TABLE questions;