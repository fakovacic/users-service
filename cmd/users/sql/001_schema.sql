-- +migrate Up

CREATE TABLE IF NOT EXISTS users (
    id              VARCHAR(50)     NOT NULL PRIMARY KEY,
    first_name      VARCHAR(255)    NOT NULL DEFAULT '',
    last_name       VARCHAR(255)    NOT NULL DEFAULT '',
    nickname        VARCHAR(255)    NOT NULL DEFAULT '',
    password        TEXT            NOT NULL DEFAULT '',
    email           TEXT            NOT NULL DEFAULT '',
    country         TEXT            NOT NULL DEFAULT '',
    created_at      TIMESTAMP,
    updated_at      TIMESTAMP,

    UNIQUE(email)
);
