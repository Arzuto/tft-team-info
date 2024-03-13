-- +migrate Up

CREATE TABLE IF NOT EXISTS class (
    id SERIAL PRIMARY KEY,
    className VARCHAR(256) NOT NULL,
    description VARCHAR(256)

);

-- +migrate Down

-- DROP TABLE IF EXISTS class;