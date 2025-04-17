-- name: CreateChat :one
INSERT INTO chat (id, user_id, content, type)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetChatByID :one
SELECT * FROM chat
WHERE id = $1;

-- name: GetChatsByUserID :many
SELECT * FROM chat
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: UpdateChatContent :one
UPDATE chat
SET content = $2
WHERE id = $1
RETURNING *;

-- name: DeleteChat :exec
DELETE FROM chat
WHERE id = $1;
