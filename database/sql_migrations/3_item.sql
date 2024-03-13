-- +migrate Up

CREATE TABLE IF NOT EXISTS item (
    id SERIAL PRIMARY KEY,
    itemName VARCHAR(256) NOT NULL,
    description VARCHAR(256),
    stats VARCHAR(256)

);

-- +migrate Down

-- DROP TABLE IF EXISTS item;