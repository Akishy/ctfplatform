-- +goose Up
CREATE TABLE IF NOT EXISTS service (
    id uuid PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS instance (
    id uuid PRIMARY KEY,
    ssh_port int NOT NULL,
    web_port int NOT NULL,
    service_id uuid REFERENCES service(id)
);

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS service;
DROP TABLE IF EXISTS instance;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
