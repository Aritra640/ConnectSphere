-- +goose Up
CREATE TABLE IF NOT EXISTS users_info (
    user_id  INTEGER PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    pimage TEXT DEFAULT '',
    pbio   TEXT DEFAULT ''
);

-- +goose Down
DROP TABLE IF EXISTS users_info;
