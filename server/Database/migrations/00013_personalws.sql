-- +goose Up
CREATE TABLE IF NOT EXISTS personal_ws (
  
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  usera INTEGER REFERENCES users(id) ON DELETE CASCADE,
  userb INTEGER REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS personal_ws;
