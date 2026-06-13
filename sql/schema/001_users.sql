-- +goose Up
CREATE TABLE (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  udpated_at TIMESTAMP NOT NULL,
  name TEXT UNIQUE NOT NULL
)
-- +goose Down
DROP TABLE users;
