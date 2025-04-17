-- +goose Up
CREATE TABLE IF NOT EXISTS chat_group (
  id  UUID  PRIMARY KEY,
  name TEXT NOT NULL,
  about TEXT NOT NULL,
  ppic  TEXT DEFAULT '',
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS chat_group;
