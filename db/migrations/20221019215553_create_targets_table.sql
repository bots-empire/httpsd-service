-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS htttp_sd

CREATE TABLE IF NOT EXISTS http_sd.targets
(
    ip_address text  NOT NULL;
    labels     jsonb NOT NULL;
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE http_sd.targerts;
-- +goose StatementEnd
