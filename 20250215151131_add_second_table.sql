-- +goose Up
-- +goose StatementBegin
CREATE TABLE second (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

ALTER TABLE team
DROP COLUMN coach;

ALTER TABLE team
ADD COLUMN coach_id INTEGER;

ALTER TABLE team
ADD CONSTRAINT club_coach
    FOREIGN KEY (coach_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE team
DROP CONSTRAINT club_coach

ALTER TABLE team DROP COLUMN coach_id;
ALTER TABLE team ADD COLUMN coach VARCHAR(255);

DROP TABLE coach;
-- +goose StatementEnd
