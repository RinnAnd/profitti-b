-- +goose Up
CREATE TABLE categories (
  id uuid NOT NULL primary key DEFAULT uuid_generate_v4 (),
  name VARCHAR(255) NOT NULL,
  user_id uuid NOT NULL references users (id)
);

ALTER TABLE expenses ADD COLUMN category_id uuid references categories (id);

-- +goose Down
ALTER TABLE expenses DROP COLUMN category_id;
DROP TABLE categories;
