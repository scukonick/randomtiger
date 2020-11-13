-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE tigers (
    id BIGSERIAL PRIMARY KEY,
    chat_id BIGINT,
    user_id BIGINT,
    stripes BIGINT,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    enlarged_at TIMESTAMPTZ DEFAULT now()
);

CREATE UNIQUE INDEX tigers_user_chat_idx ON tigers (user_id, chat_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE tigers;
-- +goose StatementEnd
