-- +goose Up
-- +goose StatementBegin
CREATE TABLE cards (
    card_id INTEGER PRIMARY KEY,
    name INTEGER,
    rareness VARCHAR(100),
    class VARCHAR(100),
    trait_skill VARCHAR(255),
    major_skill VARCHAR(255),
    nft VARCHAR(100),
);

CREATE TABLE user_card_relation (
    user_id INTEGER NOT NULL,
    card_id INTEGER NOT NULL REFERENCES cards(card_id),
    upgrade_level INTEGER NOT NULL DEFAULT 0,
    PRIMARY KEY (user_id, card_id)
);

CREATE TABLE descriptions (
    id VARCHAR(100) PRIMARY KEY,
    name_id INTEGER,
    description_id INTEGER,
    base_params JSONB,
    upgrade_step JSONB
);

CREATE TABLE traits (
    id VARCHAR(100) PRIMARY KEY,
    name_id INTEGER,
    description_id INTEGER,
    params JSONB
);

CREATE TABLE melee (
    id VARCHAR(100) PRIMARY KEY,
    name_id INTEGER,
    description_id INTEGER,
    params JSONB,
    upgrade_step JSONB
);

CREATE TABLE translations (
    text_id INTEGER PRIMARY KEY,
    ru TEXT,
    en TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS melee;
DROP TABLE IF EXISTS traits;
DROP TABLE IF EXISTS descriptions;
DROP TABLE IF EXISTS user_card_relation;
DROP TABLE IF EXISTS cards;
DROP TABLE IF EXISTS translations;
-- +goose StatementEnd
