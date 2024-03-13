-- +migrate Up

CREATE TABLE IF NOT EXISTS character (
    id SERIAL PRIMARY KEY,
    characterName VARCHAR(256) NOT NULL,
    skill VARCHAR(256),
    originIDs INT[],
    classIDs INT[],
    itemIDs INT[]
);

-- +migrate Down

-- DROP TABLE IF EXISTS character;