-- +goose Up 
CREATE TABLE IF NOT EXISTS friends (
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
  friend_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (user_id, friend_id)
);

-- +goose Down 
DROP TABLE IF EXISTS friends;
