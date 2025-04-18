-- name: CreateGroupMessage :one 
INSERT INTO chat_group_message (
  chat_id, chat_group_id, sender_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetGroupMessages :many
SELECT c.*, gm.send_at, gm.sender_id
FROM chat_group_message gm
JOIN chat c ON c.id = gm.chat_id
WHERE gm.chat_group_id = $1
ORDER BY gm.send_at; 

-- name: GetPaginatedGroupMessages :many 
SELECT c.*, gm.send_at, gm.sender_id
FROM chat_group_message gm
JOIN chat c ON c.id = gm.chat_id
WHERE gm.chat_group_id = $1
ORDER BY gm.send_at DESC
LIMIT $2 OFFSET $3; 

-- name: DeleteGroupMessage :exec
DELETE FROM chat_group_message
WHERE chat_id = $1; 

-- name: UpdateGroupMessageContent :exec
UPDATE chat
SET content = $2
WHERE id = $1;

-- name: GetNewGroupMessages :many
SELECT c.*, gm.send_at, gm.sender_id
FROM chat_group_message gm
JOIN chat c ON c.id = gm.chat_id
WHERE gm.chat_group_id = $1 AND gm.send_at > $2
ORDER BY gm.send_at;
