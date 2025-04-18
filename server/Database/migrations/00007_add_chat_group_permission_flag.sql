-- +goose Up
ALTER TABLE IF EXISTS chat_group 
ADD COLUMN required_permission BOOLEAN NOT NULL DEFAULT FALSE;

-- +goose Down
ALTER TABLE IF EXISTS chat_group 
DROP COMLUMN IF EXISTS required_permission;
