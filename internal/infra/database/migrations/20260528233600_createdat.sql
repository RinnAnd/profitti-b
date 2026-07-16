-- +goose Up
ALTER TABLE expenses ADD COLUMN created_at TIMESTAMPTZ NOT NULL DEFAULT NOW();

-- +goose Down
ALTER TABLE expenses DROP COLUMN created_at;
