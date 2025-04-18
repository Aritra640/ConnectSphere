-- name: CreateGroup :one
INSERT INTO chat_group (id, name, about, ppic , required_permission)
VALUES ($1, $2, $3, $4 , $5)
RETURNING *;

-- name: GetGroupByID :one
SELECT * FROM chat_group
WHERE id = $1;

-- name: ListGroups :many
SELECT * FROM chat_group
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: UpdateGroupInfo :one
UPDATE chat_group
SET name = $2,
    about = $3,
    ppic = $4,
    required_permission = $5
WHERE id = $1
RETURNING *;

-- name: DeleteGroup :exec
DELETE FROM chat_group
WHERE id = $1;

