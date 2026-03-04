-- +goose Up
CREATE TABLE users (
  id uuid NOT NULL primary key,
  username VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  profile VARCHAR(255)
);

CREATE TABLE currencies (
  id uuid NOT NULL primary key,
  name VARCHAR(255) NOT NULL
);

CREATE TABLE financials (
  id uuid NOT NULL primary key,
  user_id uuid NOT NULL references users (id),
  currency_id uuid NOT NULL references currencies (id)
);

CREATE TABLE recurrence_type (
  id uuid NOT NULL primary key,
  type VARCHAR(255) NOT NULL
);

CREATE TABLE partnership (
  id uuid NOT NULL primary key,
  users jsonb NOT NULL,
  currency_id uuid NOT NULL references currencies (id)
);

CREATE TABLE expenses (
  id uuid NOT NULL primary key,
  financial_id uuid references financials (id),
  partnership_id uuid references partnership (id),
  name VARCHAR(255) NOT NULL,
  description VARCHAR(255),
  amount DECIMAL,
  expense_recurrence uuid references recurrence_type (id),
  expiration_date DATE,
  currency_id uuid NOT NULL references currencies (id)
);

CREATE TABLE shared_expense (
  id uuid NOT NULL primary key,
  expense_id uuid NOT NULL references expenses (id),
  user_id uuid NOT NULL references users (id),
  percentage DECIMAL NOT NULL,
  total_amount DECIMAL,
  currency_id uuid NOT NULL references currencies (id)
);

CREATE TABLE income (
  id uuid NOT NULL primary key,
  amount DECIMAL,
  currency_id uuid NOT NULL references currencies (id),
  recurrence uuid references recurrence_type (id),
  user_id uuid NOT NULL references users (id)
);

-- +goose Down
DROP TABLE shared_expense;

DROP TABLE expenses;

DROP TABLE financials;

DROP TABLE partnership;

DROP TABLE income;

DROP TABLE users;

DROP TABLE currencies;

DROP TABLE recurrence_type;
