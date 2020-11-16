-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE tigers ADD COLUMN username TEXT DEFAULT '' NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE tigers DROP COLUMN username;
-- +goose StatementEnd
