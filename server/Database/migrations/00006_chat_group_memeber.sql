-- +goose Up
CREATE TABLE IF NOT EXISTS chat_group_member (
  group_id UUID REFERENCES chat_group(id) ON DELETE CASCADE,
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
  is_admin BOOLEAN DEFAULT FALSE,
  joined_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (group_id, user_id)
);

-- +goose Down
DROP TABLE IF EXISTS chat_group_member;
