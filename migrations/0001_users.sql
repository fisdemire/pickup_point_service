-- +goose Up
CREATE TABLE users (
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email         TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    role          TEXT NOT NULL CHECK (role IN ('employee', 'moderator')),
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now()
);
-- +goose Down
DROP TABLE users;
