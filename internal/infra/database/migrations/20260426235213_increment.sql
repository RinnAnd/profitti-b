-- +goose Up
CREATE SEQUENCE currencies_id_seq START WITH 1;
ALTER TABLE currencies ALTER COLUMN id SET DEFAULT nextval('currencies_id_seq');
ALTER SEQUENCE currencies_id_seq OWNED BY currencies.id;

-- +goose Down
ALTER TABLE currencies ALTER COLUMN id DROP DEFAULT;
DROP SEQUENCE currencies_id_seq;
