-- +migrate Up

CREATE TABLE IF NOT EXISTS recommendation (
    id SERIAL PRIMARY KEY,
    teamName VARCHAR(256) NOT NULL,
    originIDs INT[],
    classIDs INT[],
    characterIDs INT[],
    tier VARCHAR(5),
    difficulty VARCHAR(256)
);

-- +migrate Down

-- DROP TABLE IF EXISTS recommendation;