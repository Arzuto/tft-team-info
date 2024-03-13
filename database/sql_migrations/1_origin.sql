-- +migrate Up

CREATE TABLE IF NOT EXISTS origin (
    id SERIAL PRIMARY KEY,
    originName VARCHAR(256) NOT NULL,
    description VARCHAR(256)

);

-- +migrate Down

-- DROP TABLE IF EXISTS origin;