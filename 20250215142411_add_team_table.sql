-- +goose Up
-- +goose StatementBegin
CREATE TABLE team(
    id SERIAL PRIMARY KEY,
    names VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    coach VARCHAR(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
