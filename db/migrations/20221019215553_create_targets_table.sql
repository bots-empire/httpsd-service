-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS http_sd;

CREATE TABLE IF NOT EXISTS http_sd.targets
(
    ip_address text,
    labels     jsonb
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA http_sd CASCADE;
-- +goose StatementEnd
