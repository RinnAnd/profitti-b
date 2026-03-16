-- +goose Up
CREATE UNIQUE INDEX idx_unique_partnership_users ON partnership (users);

CREATE UNIQUE INDEX idx_unique_username ON users (username);

CREATE UNIQUE INDEX idx_unique_email ON users (email);

-- +goose Down
DROP INDEX idx_unique_partnership_users;

DROP INDEX idx_unique_username;

DROP INDEX idx_unique_email;
