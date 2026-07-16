-- +goose Up
CREATE TABLE friendships (
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
    requester uuid NOT NULL references users (id),
    "to" uuid NOT NULL references users (id),
    status VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE friendships;
