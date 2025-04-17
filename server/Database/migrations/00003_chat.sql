-- +goose Up
-- +goose StatementBegin
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'chat_type') THEN
        CREATE TYPE chat_type AS ENUM ('text', 'image', 'emoji', 'pdf', 'video');
    END IF;
END $$;
-- +goose StatementEnd

CREATE TABLE IF NOT EXISTS chat (
  id UUID PRIMARY KEY,
  content TEXT NOT NULL,
  type    chat_type NOT NULL DEFAULT 'text',
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS chat;
DROP TYPE IF EXISTS  chat_type;

