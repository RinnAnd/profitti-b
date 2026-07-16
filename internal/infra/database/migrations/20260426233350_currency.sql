-- +goose Up
ALTER TABLE financials DROP COLUMN currency_id;
ALTER TABLE partnership DROP COLUMN currency_id;
ALTER TABLE expenses DROP COLUMN currency_id;
ALTER TABLE shared_expense DROP COLUMN currency_id;
ALTER TABLE income DROP COLUMN currency_id;

TRUNCATE TABLE currencies CASCADE;
ALTER TABLE currencies DROP COLUMN id;
ALTER TABLE currencies ADD COLUMN id integer PRIMARY KEY;

ALTER TABLE financials ADD COLUMN currency_id integer NOT NULL references currencies (id);

ALTER TABLE partnership ADD COLUMN currency_id integer NOT NULL references currencies (id);

ALTER TABLE expenses ADD COLUMN currency_id integer NOT NULL references currencies (id);

ALTER TABLE shared_expense ADD COLUMN currency_id integer NOT NULL references currencies (id);

ALTER TABLE income ADD COLUMN currency_id integer NOT NULL references currencies (id);

-- +goose Down
SELECT 'down SQL query';
