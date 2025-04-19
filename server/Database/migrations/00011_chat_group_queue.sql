-- +goose Up
CREATE TABLE IF NOT EXISTS chat_group_queue (
  user_id   INTEGER REFERENCES users(id) ON DELETE CASCADE,
  chat_group_id UUID REFERENCES chat_group(id) ON DELETE CASCADE,
  is_accepted BOOLEAN DEFAULT FALSE,
  requested_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (user_id, chat_group_id)
);

-- +goose Down
DROP TABLE IF EXISTS chat_group_queue;
