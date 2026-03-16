-- +goose Up
ALTER TABLE expenses
  ALTER COLUMN financial_id DROP NOT NULL,
  ALTER COLUMN financial_id SET DEFAULT NULL,
  ALTER COLUMN partnership_id DROP NOT NULL,
  ALTER COLUMN partnership_id SET DEFAULT NULL;

-- +goose Down
ALTER TABLE expenses
  ALTER COLUMN financial_id SET NOT NULL,
  ALTER COLUMN financial_id DROP DEFAULT,
  ALTER COLUMN partnership_id SET NOT NULL,
  ALTER COLUMN partnership_id DROP DEFAULT;
