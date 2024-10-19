-- +goose Up
ALTER TABLE IF EXISTS users ADD COLUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT (
    encode(sha256(random()::text::bytea), 'hex')
);

-- +goose Down
ALTER TABLE IF EXISTS users DROP COLUMN api_key;
