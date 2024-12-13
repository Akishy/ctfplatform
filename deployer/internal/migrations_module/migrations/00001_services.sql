-- +goose Up
CREATE TABLE IF NOT EXISTS service (
    id uuid PRIMARY KEY,
    ssh_port INTEGER NOT NULL,
    web_port INTEGER NOT NULL
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS service;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
