-- +goose Up
ALTER TABLE test_sessions 
DROP CONSTRAINT test_sessions_template_id_fkey,
ADD CONSTRAINT test_sessions_template_id_fkey 
    FOREIGN KEY (template_id) REFERENCES test_templates(id) ON DELETE SET NULL;

-- Также нужно убрать NOT NULL, чтобы можно было поставить NULL
ALTER TABLE test_sessions ALTER COLUMN template_id DROP NOT NULL;

-- +goose Down
ALTER TABLE test_sessions ALTER COLUMN template_id SET NOT NULL;
ALTER TABLE test_sessions 
DROP CONSTRAINT test_sessions_template_id_fkey,
ADD CONSTRAINT test_sessions_template_id_fkey 
    FOREIGN KEY (template_id) REFERENCES test_templates(id);
