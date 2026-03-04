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

-- +goose Down
DROP TABLE financials;

DROP TABLE users;

DROP TABLE currencies;
